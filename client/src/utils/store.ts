import { createStore } from "solid-js/store"
import { createResource, createSignal, ResourceReturn, Signal } from "solid-js"
import { Pixel, Theme } from "./models"
import { api } from "./api"

async function fetchPixels(): Promise<Map<number, Signal<string>>> {
  const res = await api.pixels.get()
  if (!res.ok) {
    throw res.error
  }

  const pixels = new Map<number, Signal<string>>()
  
  res.data.forEach(({ id, color }) => {
    pixels.set(id, createSignal(color))
  })

  return pixels
}

export type AppState = {
  theme: Theme
  pixelsResource: ResourceReturn<Map<number, Signal<string>>>
  selectedColor: string
}

const [store, setStore] = createStore<AppState>({
  theme: null,
  pixelsResource: createResource(fetchPixels),
  selectedColor: "#ff0000",
})

async function setPixel(pixel: Pixel) {
  const [pixels] = store.pixelsResource
  const signal = pixels()?.get(pixel.id)
  if (!signal) return

  const [_, setValue] = signal
  setValue(pixel.color)

  await api.pixels.post({
    body: pixel
  })
}

function setTheme(theme: Theme) {
  setStore({
    theme,
  })
}

function setColor(color: string) {
  setStore({
    selectedColor: color
  })
}

export { store, setPixel, setTheme, setColor }
