import { jex } from "@olli/jex";
import { PixelSchema } from "./models";

export const api = jex({
  baseUrl: "",
  endpoints: {
    pixels: {
      get: {
        data: PixelSchema.array()
      },
      post: {
        body: PixelSchema
      }
    }
  }
})
