export interface Tileset {
  name: string
  columns: number
  grid: {
    width: number
    height: number
    orientation: string
  }
  image: string
  imageheight: number
  imagewidth: number
  margin: number
  spacing: number
  tilecount: number
  tiledversion: string
  tilewidth: number
  tileheight: number
  type: string
  version: number
  tiles: TileDefinition[]
}

interface TileDefinition {
  id: number
  type: string
}

export interface Tilemap {
  compressionlevel: number
  height: number
  width: number
  orientation: string
  renderorder: string
  staggeraxis: string
  staggerindex: string
  tiledversion: string
  version: number
  tileheight: number
  tilewidth: number
  tilesets: TilemapTileset[]
  layers: TilemapLayer[]
  type: string
}

interface TilemapTileset {
  firstgid: number
  source: string
}

interface TilemapLayer {
  id: number
  name: string
  type: string
  data: number[]
  height: number
  width: number
  x: number
  y: number
  visible: boolean
  opacity: number
}