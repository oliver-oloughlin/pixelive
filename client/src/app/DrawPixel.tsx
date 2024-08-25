import { Signal } from "solid-js";
import { setPixel, store } from "../utils/store"
import { apiWs } from "../utils/api_ws";
import { Pixel } from "../utils/models";

export function DrawPixel({ id, signal }: { id: number, signal: Signal<string> }) {
  const [color] = signal

  const setter = () => {
    const pixel: Pixel = { id, color: store.selectedColor }
    setPixel(pixel)
    apiWs.setPixel(pixel)
  }

  return (
    <div 
      onPointerDown={setter}
      class="hover:backdrop-brightness-75 hover:invert-[50%] aspect-square"
      style={{ "background-color": color() }}
    />
  )
}
