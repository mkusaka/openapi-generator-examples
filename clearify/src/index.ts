import { petApi } from "../api/config"
(async () => {
  await petApi.addPet({
    pet: {
      name: "test",
      photoUrls: ["https://example.example"]
    }
  })
})()
