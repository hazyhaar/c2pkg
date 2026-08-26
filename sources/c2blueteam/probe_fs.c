/*
 * probe_fs.c — Module de surveillance de fichiers fanotify avec protection anti-deadlock.
 */

#define _GNU_SOURCE
#include "c2blueteam.h"
#include "ring_buffer.h"
#include <sys/fanotify.h>
#include <fcntl.h>
#include <unistd.h>
#include <string.h>
#include <stdio.h>

int c2bt_probe_fs_init(int *out_fan_fd, const char *protected_dir) {
    if (!out_fan_fd) return -1;
    
    /* Initialisation fanotify non-bloquant par défaut (mode notification) */
    int fd = fanotify_init(FAN_CLASS_NOTIF | FAN_NONBLOCK, O_RDONLY | O_LARGEFILE | O_CLOEXEC);
    if (fd < 0) {
        return -1;
    }
    
    if (protected_dir && protected_dir[0] != '\0') {
        fanotify_mark(fd, FAN_MARK_ADD, FAN_OPEN | FAN_MODIFY | FAN_CLOSE_WRITE, AT_FDCWD, protected_dir);
    }
    
    *out_fan_fd = fd;
    return 0;
}

int c2bt_probe_fs_init_perm(int *out_fan_fd, const char *protected_dir) {
    if (!out_fan_fd) return -1;
    
    /* Initialisation fanotify synchrone (mode permission avec interception préalable) */
    int fd = fanotify_init(FAN_CLASS_CONTENT | FAN_NONBLOCK, O_RDWR | O_LARGEFILE | O_CLOEXEC);
    if (fd < 0) {
        return -1;
    }
    
    if (protected_dir && protected_dir[0] != '\0') {
        fanotify_mark(fd, FAN_MARK_ADD, FAN_OPEN_PERM | FAN_ACCESS_PERM | FAN_OPEN | FAN_MODIFY, AT_FDCWD, protected_dir);
    }
    
    *out_fan_fd = fd;
    return 0;
}

int c2bt_probe_fs_verdict(int fanotify_fd, int fd_event, uint32_t flags, const c2bt_config_t *cfg) {
    if (fanotify_fd < 0 || fd_event < 0) {
        return -1;
    }
    
    struct fanotify_response resp;
    memset(&resp, 0, sizeof(resp));
    resp.fd = fd_event;
    
    /* Veto actif : émission de FAN_DENY si le drapeau BLOCKED est levé et enforce_mode activé */
    if (cfg && (cfg->enforce_mode == C2BT_MODE_ACTIVE) && (flags & C2BT_FLAG_BLOCKED)) {
        resp.response = FAN_DENY;
    } else {
        resp.response = FAN_ALLOW;
    }
    
    ssize_t w = write(fanotify_fd, &resp, sizeof(resp));
    if (w != (ssize_t)sizeof(resp)) {
        return -1;
    }
    return 0;
}

int c2bt_probe_fs_poll(int fan_fd, probe_channel_t *ch) {
    if (fan_fd < 0 || !ch) return 0;
    
    char buf[4096];
    ssize_t len = read(fan_fd, buf, sizeof(buf));
    if (len <= 0) return 0;
    
    const struct fanotify_event_metadata *metadata = (const struct fanotify_event_metadata *)buf;
    int count = 0;
    
    while (FAN_EVENT_OK(metadata, len)) {
        if (metadata->vers >= FANOTIFY_METADATA_VERSION) {
            probe_event_t ev = {0};
            ev.ts_ns = c2bt_time_ns_raw();
            ev.pid = (uint32_t)metadata->pid;
            ev.subsystem = C2BT_SUB_FS;
            
            if (metadata->mask & (FAN_OPEN_PERM | FAN_ACCESS_PERM)) {
                ev.action = (metadata->mask & FAN_ACCESS_PERM) ? C2BT_ACT_WRITE : C2BT_ACT_OPEN;
            } else {
                ev.action = (metadata->mask & FAN_MODIFY) ? C2BT_ACT_WRITE : C2BT_ACT_OPEN;
            }
            ev.src = (uint64_t)metadata->fd;
            
            char fd_path[64];
            snprintf(fd_path, sizeof(fd_path), "/proc/self/fd/%d", metadata->fd);
            ssize_t r = readlink(fd_path, (char *)ev.payload, sizeof(ev.payload) - 1);
            if (r > 0) {
                ev.payload[r] = '\0';
            }
            
            c2bt_channel_write(ch, &ev);
            count++;
            
            /* Réponse synchrone au kernel si événement de type permission */
            if (metadata->mask & (FAN_OPEN_PERM | FAN_ACCESS_PERM)) {
                c2bt_probe_fs_verdict(fan_fd, metadata->fd, ev.flags, NULL);
            }
            
            if (metadata->fd >= 0) {
                close(metadata->fd);
            }
        }
        metadata = FAN_EVENT_NEXT(metadata, len);
    }
    
    return count;
}
