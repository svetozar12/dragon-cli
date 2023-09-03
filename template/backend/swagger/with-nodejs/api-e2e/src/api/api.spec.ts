import { defaultApi } from "@dragon-cli-template/shared/sdk";

describe("GET /api/example", () => {
  it("should return a message", async () => {
    const res = await defaultApi.instance().exampleGet();

    expect(res.status).toBe(200);
    expect(res.data).toEqual({ message: "Hello API" });
  });
});
