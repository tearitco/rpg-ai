import React, { useContext, useEffect, useRef } from 'react';
import { Context, initialState, Pos, setTile, State } from './State';

class CanvasRenderer {
  mouse?: Pos = undefined
  mouseDown: boolean = false
  drag?: {
    start: Pos
    end: Pos
  }
  size: number = 25
  lastTime: number = 0
  requestID?: number
  appState: State = initialState
  canvas: HTMLCanvasElement
  ctx: CanvasRenderingContext2D

  constructor(canvas: HTMLCanvasElement) {
    const ctx = canvas.getContext('2d');
    if (ctx === null) {
      throw new Error('canvas has null rendering context');
    }
    this.canvas = canvas;
    this.ctx = ctx
    canvas.addEventListener('mousedown', this.onMouseDown);
    canvas.addEventListener('mouseup', this.onMouseUp);
    canvas.addEventListener('mousemove', this.onMouseMove);
    this.requestID = requestAnimationFrame(this.render);
  }

  onMouseDown = () => {
    this.mouseDown = true;
    if (this.mouse) {
      this.drag = {
        start: { ...this.mouse },
        end: { ...this.mouse },
      }
    }
  }

  onMouseUp = () => {
    this.mouseDown = false;
    if (this.drag) {
      this.drag = undefined;
    }
  }

  onMouseMove = (event: MouseEvent) => {
    this.mouse = {
      x: event.clientX,
      y: event.clientY,
    };
    if (this.mouse && this.mouseDown && this.drag) {
      this.drag.end = { ...this.mouse };
    }
    if (this.appState.tools.brush.selected && this.mouse && this.mouseDown) {
      const x = Math.floor(this.mouse.x / this.size);
      const y = Math.floor(this.mouse.y / this.size);
      if (!this.appState.map[x] || !this.appState.map[x][y]) {
        setTile(this.appState, { x, y });
      }
    }
  }

  renderTextCenter(text: string, font: string) {
    this.ctx.font = font;
    const m = this.ctx.measureText(text);
    let h = m.actualBoundingBoxAscent + m.actualBoundingBoxDescent;
    this.ctx.fillText(text, -m.width / 2, h);
  }

  renderFPS(time: number) {
    const fps = (1000 / (time - this.lastTime)).toFixed(0);
    this.ctx.translate(this.canvas.width - 18, 12);
    this.renderTextCenter(fps, "18px Roboto Mono, monospace");
  }

  clearScreen() {
    const { canvas, ctx } = this;
    ctx.setTransform(1, 0, 0, 1, 0, 0);
    ctx.fillStyle = '#D9D2BF';
    ctx.fillRect(0, 0, canvas.width, canvas.height);
  }

  drawGrid() {
    const { canvas, ctx } = this;
    for (let x = 0; x < canvas.width; x += this.size) {
      for (let y = 0; y < canvas.height; y += this.size) {
        ctx.lineWidth = 0.1;
        ctx.strokeStyle = '#000';
        ctx.beginPath();
        ctx.rect(x, y, this.size, this.size);
        ctx.stroke();
      }
    }
  }

  drawMousePos() {
    const { ctx } = this;
    if (this.mouse) {
      ctx.beginPath();
      ctx.fillStyle = '#000';
      ctx.arc(
        Math.round(this.mouse.x / this.size) * this.size,
        Math.round(this.mouse.y / this.size) * this.size,
        2, 0, 2 * Math.PI);
      ctx.fill();
    }
  }

  drawDrag() {
    const { ctx } = this;
    if (this.drag) {
      const x1 = Math.round(this.drag.start.x / this.size) * this.size;
      const y1 = Math.round(this.drag.start.y / this.size) * this.size;
      const x2 = Math.round(this.drag.end.x / this.size) * this.size;
      const y2 = Math.round(this.drag.end.y / this.size) * this.size;
      ctx.beginPath();
      ctx.fillStyle = '#000';
      ctx.arc(x1, y1, 2, 0, 2 * Math.PI);
      ctx.fill();
      ctx.beginPath();
      ctx.arc(x2, y2, 2, 0, 2 * Math.PI);
      ctx.fill();

      ctx.beginPath();
      ctx.strokeStyle = '#2f5574';
      ctx.lineWidth = 2;
      ctx.moveTo(x1, y1);
      ctx.lineTo(x1, y2);
      ctx.lineTo(x2, y2);
      ctx.lineTo(x2, y1);
      ctx.lineTo(x1, y1);
      ctx.stroke();

      const w = Math.abs(x1 - x2) / this.size;
      const h = Math.abs(y1 - y2) / this.size;
      ctx.translate(x1 + (x2 - x1) / 2, y1 - 14);
      this.renderTextCenter(`${w} x ${h}`, "14px Roboto, sans-serif");
    }
  }

  drawTile() {
    const { ctx } = this;
    ctx.beginPath();
    ctx.fillStyle = '#F1ECE0';
    ctx.fillRect(0, 0, this.size, this.size);
  }

  drawHoverTile() {
    const { ctx } = this;
    if (this.mouse) {
      ctx.translate(
        Math.floor(this.mouse.x / this.size) * this.size,
        Math.floor(this.mouse.y / this.size) * this.size);
      this.drawTile();
    }
  }

  drawMap() {
    const { ctx, appState } = this;
    Object.entries(appState.map).forEach(([x, col]) => Object.entries(col).forEach(([y, occupied]) => {
      if (occupied) {
        ctx.save();
        ctx.translate(parseInt(x) * this.size, parseInt(y) * this.size);
        this.drawTile();
        ctx.restore();
      }
    }));
  }

  render = (time: number) => {
    const { ctx, appState } = this;

    ctx.save();
    this.clearScreen();
    ctx.restore();

    ctx.save();
    this.renderFPS(time);
    ctx.restore();

    ctx.save();
    this.drawGrid();
    ctx.restore();

    if (appState.tools.brush.selected) {
      ctx.save();
      this.drawHoverTile();
      ctx.restore();
    } else {
      ctx.save();
      this.drawMousePos();
      ctx.restore();
    }

    if (appState.tools.box.selected || appState.tools.circle.selected) {
      ctx.save();
      this.drawDrag();
      ctx.restore();
    }

    ctx.save();
    this.drawMap();
    ctx.restore();

    this.lastTime = time;
    this.requestID = requestAnimationFrame(this.render);
  }
}

export default function Canvas() {
  const canvasRef = useRef<HTMLCanvasElement>(null);
  const appState = useContext(Context);
  useEffect(() => {
    if (canvasRef.current === null) return;
    const canvas = canvasRef.current;
    canvas.width = canvas.parentElement?.offsetWidth || canvas.width;
    canvas.height = canvas.parentElement?.offsetHeight || canvas.height;
    if ((canvas as any).renderer === undefined) {
      (canvas as any).renderer = new CanvasRenderer(canvas);
    }
    (canvas as any).renderer.appState = appState;
  }, [canvasRef, appState]);
  return <canvas ref={canvasRef} />;
}
