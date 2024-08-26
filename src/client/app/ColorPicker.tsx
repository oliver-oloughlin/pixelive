import { store, setColor } from "../utils/store";

export function ColorPicker() {
  return (
    <label class="flex gap-2 items-center font-semibold text-2xl">
      Color:
      <input
        class="h-10 w-20"
        value={store.selectedColor}
        type="color"
        onInput={e => setColor(e.target.value)}
      />
    </label>
  )
}
