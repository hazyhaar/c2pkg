/*
 * probe_proc.c — Normalisation et capture des événements de cycle de vie des processus.
 */

#define _GNU_SOURCE
#include "c2blueteam.h"
#include "ring_buffer.h"
#include <fcntl.h>
#include <unistd.h>
#include <string.h>
#include <stdio.h>

int c2bt_probe_proc_read_cmdline(uint32_t pid, char *out_cmdline, size_t max_len) {
    if (!out_cmdline || max_len == 0) return -1;
    
    char path[64];
    snprintf(path, sizeof(path), "/proc/%u/cmdline", (unsigned int)pid);
    
    int fd = open(path, O_RDONLY | O_NOFOLLOW);
    if (fd < 0) {
        out_cmdline[0] = '\0';
        return -1;
    }
    
    ssize_t n = read(fd, out_cmdline, max_len - 1);
    close(fd);
    
    if (n <= 0) {
        out_cmdline[0] = '\0';
        return 0;
    }
    
    for (ssize_t i = 0; i < n; i++) {
        if (out_cmdline[i] == '\0' && i + 1 < n) {
            out_cmdline[i] = ' ';
        }
    }
    out_cmdline[n] = '\0';
    return (int)n;
}

int c2bt_probe_proc_package_exec(probe_channel_t *ch, uint32_t pid, uint32_t ppid, const char *comm, const char *cmdline) {
    if (!ch) return -1;
    
    probe_event_t ev = {0};
    ev.ts_ns = c2bt_time_ns_raw();
    ev.pid = pid;
    ev.tid = pid;
    ev.subsystem = C2BT_SUB_PROC;
    ev.action = C2BT_ACT_EXEC;
    ev.src = (uint64_t)ppid;
    
    if (comm && comm[0] != '\0') {
        snprintf((char *)ev.payload, sizeof(ev.payload), "%.95s", comm);
    } else if (cmdline && cmdline[0] != '\0') {
        snprintf((char *)ev.payload, sizeof(ev.payload), "%.95s", cmdline);
    }
    
    return c2bt_channel_write(ch, &ev);
}

int c2bt_probe_proc_package_fork(probe_channel_t *ch, uint32_t parent_pid, uint32_t child_pid) {
    if (!ch) return -1;
    
    probe_event_t ev = {0};
    ev.ts_ns = c2bt_time_ns_raw();
    ev.pid = child_pid;
    ev.subsystem = C2BT_SUB_PROC;
    ev.action = C2BT_ACT_EXEC;
    ev.src = (uint64_t)parent_pid;
    snprintf((char *)ev.payload, sizeof(ev.payload), "fork(ppid=%u, child=%u)", parent_pid, child_pid);
    
    return c2bt_channel_write(ch, &ev);
}
