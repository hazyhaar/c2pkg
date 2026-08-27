#ifndef C2RASTER_H
#define C2RASTER_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

/*
 * rasterize_bezier_quad calcule analytiquement la couverture anti-crénelée (0-255)
 * d'une courbe de Bézier quadratique définie par les points de contrôle (x0, y0),
 * (x1, y1), (x2, y2) dans un masque d'octets 8-bit de dimensions w x h.
 * Aucune allocation dynamique.
 */
void rasterize_bezier_quad(float x0, float y0, float x1, float y1, float x2, float y2, uint8_t *coverage, int w, int h);

/*
 * rasterize_rounded_rect trace un rectangle à coins arrondis analytiquement anti-crénelé
 * dans le tampon de pixels 32-bit framebuffer (largeur/stride).
 * Aucune allocation dynamique.
 */
void rasterize_rounded_rect(uint32_t *framebuffer, int stride, int x, int y, int w, int h, int radius, uint32_t color);

#ifdef __cplusplus
}
#endif

#endif /* C2RASTER_H */
