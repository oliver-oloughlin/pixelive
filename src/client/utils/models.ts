import { z } from "zod"

export const PixelSchema = z.object({
  id: z.number(),
  color: z.string(),
})

export const ThemeSchema = z.enum(["light", "dark"]).or(z.null())

export type Pixel = z.infer<typeof PixelSchema>

export type Theme = z.infer<typeof ThemeSchema>
