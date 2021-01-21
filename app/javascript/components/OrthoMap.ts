import Phaser from 'phaser';
import EasyStar from 'easystarjs';

import consumer from '../channels/consumer';

export default class OrthoMap extends Phaser.Scene {
  tiledMap: any;
  mapData: Phaser.Tilemaps.MapData;
  map: Phaser.Tilemaps.Tilemap;

  cursors: Phaser.Types.Input.Keyboard.CursorKeys
  controls: Phaser.Cameras.Controls.SmoothedKeyControl
  objects: Phaser.GameObjects.Group[] = []
  player: Phaser.GameObjects.Sprite
  movementKeys: {
    up: Phaser.Input.Keyboard.Key
    down: Phaser.Input.Keyboard.Key
    left: Phaser.Input.Keyboard.Key
    right: Phaser.Input.Keyboard.Key
    shift: Phaser.Input.Keyboard.Key
  }
  pathFinder: EasyStar.js
  marker: Phaser.GameObjects.Graphics
  wallLayer: Phaser.Tilemaps.LayerData
  npcs: any[]

  init(args: { tilemap: any, character: { name: string, sprite: string } }) {
    this.tiledMap = args.tilemap;
    this.mapData = Phaser.Tilemaps.Parsers.Tiled.ParseJSONTiled(args.tilemap.name, args.tilemap, false);
    this.textures.addBase64('character', args.character.sprite);
  }

  preload() {
    this.mapData.tilesets.forEach(tileset => {
      const tilesetDef = this.tiledMap.tilesets.find((t: any) => t.name == tileset.name);
      this.load.spritesheet(
        tileset.name,
        tilesetDef.image,
        { frameWidth: tilesetDef.tilewidth, frameHeight: tilesetDef.tileheight });
    });
  }

  create() {
    const controlConfig = {
      camera: this.cameras.main,
      acceleration: 0.5,
      drag: 0.01,
      maxSpeed: 2.0,
      zoomIn: this.input.keyboard.addKey(Phaser.Input.Keyboard.KeyCodes.Q),
      zoomOut: this.input.keyboard.addKey(Phaser.Input.Keyboard.KeyCodes.E),
    };
    this.controls = new Phaser.Cameras.Controls.SmoothedKeyControl(controlConfig);
    this.cameras.main.setBackgroundColor('#ddd');
    this.cameras.main.setZoom(2, 2);
    this.map = this.make.tilemap(this.mapData);
    const tilesets = this.mapData.tilesets.map(tileset => {
      return this.map.addTilesetImage(
        tileset.name,
        tileset.name,
        tileset.tileWidth, tileset.tileHeight,
        tileset.tileMargin, tileset.tileSpacing,
        tileset.firstgid);
    });
    const layers = this.mapData.layers;
    if (layers instanceof Array) {
      this.map.layers = layers;
      layers.forEach((layer, i) => {
        this.map.createLayer(layer.name, tilesets);
        if (i === 1) {
          this.player = this.add.sprite(0, 0, "character", 0);
        }
      });
    }
    this.movementKeys = {
      up: this.input.keyboard.addKey(Phaser.Input.Keyboard.KeyCodes.UP),
      down: this.input.keyboard.addKey(Phaser.Input.Keyboard.KeyCodes.DOWN),
      left: this.input.keyboard.addKey(Phaser.Input.Keyboard.KeyCodes.LEFT),
      right: this.input.keyboard.addKey(Phaser.Input.Keyboard.KeyCodes.RIGHT),
      shift: this.input.keyboard.addKey(Phaser.Input.Keyboard.KeyCodes.SHIFT),
    };
    Object.values(this.movementKeys).forEach(key => key.setEmitOnRepeat(true));
    this.movementKeys.up.on('down', this.moveUp);
    this.movementKeys.down.on('down', this.moveDown);
    this.movementKeys.left.on('down', this.moveLeft);
    this.movementKeys.right.on('down', this.moveRight);
    this.cameras.main.startFollow(this.player);
    (window as any).map = this.map;
    (window as any).mapData = this.mapData;
    (window as any).tiledMap = this.tiledMap;
    const spawn = this.findObject(['spawn']);
    if (spawn) {
      const x = spawn.x + Math.random() * spawn.width;
      const y = spawn.y + Math.random() * spawn.height;
      this.player.setPosition(
        Math.floor(x / this.mapData.tileWidth) * this.mapData.tileWidth + (this.mapData.tileWidth / 2),
        Math.floor(y / this.mapData.tileHeight) * this.mapData.tileHeight + (this.mapData.tileHeight / 2));
    }
    this.wallLayer = this.map.layers.find(layer => layer.name === "Walls");
    this.pathFinder = new EasyStar.js();
    const grid = [];
    for (let y = 0; y < this.map.height; y++) {
      let col = [];
      for (var x = 0; x < this.map.width; x++) {
        let floor = false;
        let collision = false;
        this.map.layers.forEach(layer => {
          if (layer.name === 'Floor' && layer.data[y][x].index !== -1) floor = true;
          if (layer.name === 'Walls' && layer.data[y][x].index !== -1) collision = true;
          if (layer.data[y][x].properties['collides']) collision = true;
        });
        const doors = (this.mapData.objects as Phaser.Types.Tilemaps.ObjectLayerConfig[]).find(objLayer => objLayer.name === 'Doors');
        const doorsHit = doors.objects.filter(obj => new Phaser.Geom.Rectangle(obj.x, obj.y, obj.width, obj.height).contains(x * this.map.tileWidth, y * this.map.tileHeight));
        if (floor && (!collision || doorsHit.length > 0)) {
          col.push(0);
        } else {
          col.push(1);
        }
      }
      grid.push(col);
    }
    this.pathFinder.setGrid(grid);
    this.pathFinder.setAcceptableTiles([0]);
    this.pathFinder.enableDiagonals();
    this.pathFinder.enableCornerCutting();
    this.input.on('pointerup', this.handleClick);
    const title = this.add.text(
      this.scale.canvas.width / 2,
      375,
      this.mapData.name,
      {
        color: '#000000',
        fontFamily: 'Georgia',
        fontSize: '24px',
        resolution: window.devicePixelRatio,
      },
    );
    title.setScrollFactor(0);
    title.setOrigin(0.5);
    this.add.tween({
      targets: title,
      duration: 5000,
      ease: 'Linear',
      delay: 1000,
      alpha: 0,
    });
    /*
    const subscription = consumer.subscriptions.create({
      channel: "TilemapChannel",
      tilemap_id: this.tiledMap.id,
    }, {
      received: (data: any) => {
        console.log(data);
      },
    });
    (window as any).subscription = subscription;
    */
    this.map.layers.forEach(layer => layer.data.forEach(row => row.forEach(tile => tile.properties['seen'] = false)));
    this.updateVisibility();

    this.marker = this.add.graphics();
    this.marker.lineStyle(1, 0xffffff, 1);
    this.marker.strokeRect(0, 0, this.map.tileWidth, this.map.tileHeight);

    (this.mapData.objects as Phaser.Types.Tilemaps.ObjectLayerConfig[]).filter(layer => layer.name === "NPCs").forEach(layer => {
      console.log(layer);
    });
  }

  handleClick = (pointer: Phaser.Input.Pointer) => {
    const worldPoint = this.input.activePointer.positionToCamera(this.cameras.main);
    if (!('x' in worldPoint && 'y' in worldPoint)) {
      return;
    }
    const toX = this.map.worldToTileX(worldPoint.x);
    const toY = this.map.worldToTileY(worldPoint.y);
    const fromX = Math.floor(this.player.x / this.map.tileWidth);
    const fromY = Math.floor(this.player.y / this.map.tileHeight);
    this.pathFinder.findPath(fromX, fromY, toX, toY, path => {
      if (path === null) {
        console.log('no path');
      } else {
        this.moveCharacter(path);
      }
    });
    this.pathFinder.calculate();
  }

  moveCharacter = (path: { x: number, y: number }[]) => {
    const tweens = [];
    for (let i = 0; i < path.length - 1; i++) {
      const ex = path[i + 1].x;
      const ey = path[i + 1].y;
      tweens.push({
        targets: this.player,
        x: { value: ex * this.map.tileWidth + this.map.tileWidth / 2, duration: 200 },
        y: { value: ey * this.map.tileHeight + this.map.tileHeight / 2, duration: 200 }
      });
    }
    this.tweens.timeline({
      tweens: tweens
    });
  }

  findObject(properties: string[]) {
    for (let layer of (this.mapData.objects as Phaser.Types.Tilemaps.ObjectLayerConfig[])) {
      for (let obj of layer.objects) {
        if (properties.every(prop => obj.properties.find(p => p.name === prop))) {
          return obj;
        }
      }
    }
  }

  findObjects(properties: string[]) {
    const matches = [];
    for (let layer of (this.mapData.objects as Phaser.Types.Tilemaps.ObjectLayerConfig[])) {
      for (let obj of layer.objects) {
        if (properties.every(prop => obj.properties.find(p => p.name === prop))) {
          matches.push(obj);
        }
      }
    }
    return matches;
  }

  updateVisibility() {
    this.map.layers.forEach(layer => layer.data.forEach(row => row.forEach(tile => {
      tile.tint = 0xffffff;
      if (!tile.properties['seen']) tile.visible = false;
      else tile.tint = 0xaaaaaa;
    })));
    for (let x = 0; x < this.map.width; x++) {
      this.updateVisibilityAlongLine(x, 0);
      this.updateVisibilityAlongLine(x, this.map.height - 1);
    }
    for (let y = 0; y < this.map.width; y++) {
      this.updateVisibilityAlongLine(0, y);
      this.updateVisibilityAlongLine(this.map.width - 1, y);
    }
  }

  updateVisibilityAlongLine(toX: number, toY: number) {
    const px = Math.floor(this.player.x / this.map.tileWidth);
    const py = Math.floor(this.player.y / this.map.tileHeight);
    const line = new Phaser.Geom.Line(px, py, toX, toY);
    const points = Phaser.Geom.Line.BresenhamPoints(line);
    for (let j = 0; j < points.length; j++) {
      const point = points[j];
      this.map.layers.forEach(layer => {
        if (layer.name.includes("Secret")) {
          // TODO: Reveal secrets
          return;
        }
        if (point.y in layer.data && point.x in layer.data[point.y]) {
          const tile = layer.data[point.y][point.x];
          tile.visible = true;
          tile.properties['seen'] = true;
          tile.tint = 0xffffff;
        }
      });
      if (point.y in this.wallLayer.data && point.x in this.wallLayer.data[point.y] && this.wallLayer.data[point.y][point.x].index >= 0) {
        return;
      }
    }
  }

  avoidCollision(oldX: number, oldY: number) {
    let collision = this.map.layers.map(layer => {
      return this.map.getTilesWithinWorldXY(this.player.x - 8, this.player.y - 8, 16, 16, null, null, layer.name);
    }).flat().some(tile => tile.index !== -1 && (tile.layer.name === 'Walls' || tile.properties['collides']));
    const doors = (this.mapData.objects as Phaser.Types.Tilemaps.ObjectLayerConfig[]).find(objLayer => objLayer.name === 'Doors');
    const doorsHit = doors.objects.filter(obj => new Phaser.Geom.Rectangle(obj.x, obj.y, obj.width, obj.height).contains(this.player.x, this.player.y));
    if (doorsHit.length > 0) {
      collision = false;
    }
    if (collision && !this.movementKeys.shift.isDown) {
      this.player.x = oldX;
      this.player.y = oldY;
    } else {
      this.updateVisibility();
    }
  }

  moveUp = () => {
    const [oldX, oldY] = [this.player.x, this.player.y];
    this.player.y -= 16
    if (this.movementKeys.left.isDown) this.player.x -= 16;
    if (this.movementKeys.right.isDown) this.player.x += 16;
    this.avoidCollision(oldX, oldY);
  }

  moveDown = () => {
    const [oldX, oldY] = [this.player.x, this.player.y];
    this.player.y += 16
    if (this.movementKeys.left.isDown) this.player.x -= 16;
    if (this.movementKeys.right.isDown) this.player.x += 16;
    this.avoidCollision(oldX, oldY);
  }

  moveLeft = () => {
    const [oldX, oldY] = [this.player.x, this.player.y];
    this.player.x -= 16;
    if (this.movementKeys.up.isDown) this.player.y -= 16;
    if (this.movementKeys.down.isDown) this.player.y += 16;
    this.avoidCollision(oldX, oldY);
  }

  moveRight = () => {
    const [oldX, oldY] = [this.player.x, this.player.y];
    this.player.x += 16;
    if (this.movementKeys.up.isDown) this.player.y -= 16;
    if (this.movementKeys.down.isDown) this.player.y += 16;
    this.avoidCollision(oldX, oldY);
  }

  update(time: number, delta: number) {
    this.controls?.update(delta);
    if (this.tweens.isTweening(this.player)) {
      this.updateVisibility();
    }

    const worldPoint = this.input.activePointer.positionToCamera(this.cameras.main);
    if ('x' in worldPoint && 'y' in worldPoint) {
      const pointerTileX = this.map.worldToTileX(worldPoint.x);
      const pointerTileY = this.map.worldToTileY(worldPoint.y);
      this.marker.x = this.map.tileToWorldX(pointerTileX);
      this.marker.y = this.map.tileToWorldY(pointerTileY);
    }
  }

}