import { createStore } from "solid-js/store"
import { createResource, createSignal, ResourceReturn, Signal } from "solid-js"
import { Pixel, Theme } from "./models"
import { api } from "./api"

function defaultPixels() {
  const pixels: Pixel[] = []
  for (let i = 0; i < 100 * 100; i++) {
    pixels.push({
      index: i,
      color: "transparent"
    })
  }
  return pixels
}

async function fetchPixels(): Promise<Pixel[]> {
  const res = await api.pixels.get()
  if (!res.ok) return defaultPixels()

  return res.data
}

export type AppState = {
  theme: Theme
  pixelsResource: ResourceReturn<Pixel[]>
  selectedColor: string
  pixels: Signal<Pixel>[]
}

const [store, setStore] = createStore<AppState>({
  theme: null,
  pixelsResource: createResource(fetchPixels),
  selectedColor: "black",
  pixels: defaultPixels().map((pixel) => createSignal(pixel))
})

async function setPixel(pixel: Pixel) {
  const signal = store.pixels[pixel.index]
  if (!signal) return

  const [_, setValue] = signal
  setValue(pixel)

  await api.pixels.post({
    body: pixel
  })
}

function setTheme(theme: Theme) {
  setStore({
    theme,
  })
}

export { store, setPixel, setTheme }
