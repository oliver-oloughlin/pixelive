import { store, setColor } from "../utils/store";

export function ColorPicker() {
  return (
    <label class="flex gap-2 font-semibold">
      Color:
      <input
        value={store.selectedColor}
        type="color"
        onInput={e => setColor(e.target.value)}
      />
    </label>
  )
}
