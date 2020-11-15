export interface Pos {
  x: number
  y: number
}

export function line(x0: number, y0: number, x1: number, y1: number): { x: number, y: number }[] {
  const dx = Math.abs(x1 - x0);
  const dy = Math.abs(y1 - y0);
  const sx = (x0 < x1) ? 1 : -1;
  const sy = (y0 < y1) ? 1 : -1;
  let err = dx - dy;
  const points = [];
  while (true) {
    points.push({ x: x0, y: y0 });
    if ((x0 === x1) && (y0 === y1)) break;
    var e2 = 2 * err;
    if (e2 > -dy) { err -= dy; x0 += sx; }
    if (e2 < dx) { err += dx; y0 += sy; }
  }
  return points;
}

export function sqDist(x0: number, y0: number, x1: number, y1: number): number {
  return (x1 - x0) * (x1 - x0) + (y1 - y0) * (y1 - y0);
}

export function dist(x0: number, y0: number, x1: number, y1: number): number {
  return Math.sqrt(sqDist(x0, y0, x1, y1));
}

export function rgbToHex(r: number, g: number, b: number): string {
  if (r < 0 || r > 255 || g < 0 || g > 255 || b < 0 || b > 255) throw new Error('Invalid color component');
  return ((r << 16) | (g << 8) | b).toString(16);
}

export function colorToIndex(r: number, g: number, b: number): number {
  return ((r << 16) | (g << 8) | b);
}

export function indexToColor(i: number): string {
  if (i < 0 || i > 255 * 255 * 255) throw new Error('Index out of range');
  return '#' + i.toString(16).padStart(6, '0');
}

export function boundingRect(from: Pos, to: Pos, closest: number = 1) {
  let x1 = Math.min(from.x, to.x);
  let y1 = Math.min(from.y, to.y);
  let x2 = Math.max(from.x, to.x);
  let y2 = Math.max(from.y, to.y);
  x1 = Math.round(x1 / closest) * closest;
  y1 = Math.round(y1 / closest) * closest;
  x2 = Math.round(x2 / closest) * closest;
  y2 = Math.round(y2 / closest) * closest;
  return [{ x: x1, y: y1 }, { x: x2, y: y2 }];
}