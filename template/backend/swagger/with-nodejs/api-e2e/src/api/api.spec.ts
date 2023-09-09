import { defaultApi } from "@dragon-cli-template/shared/sdk";

describe("GET /api/example", () => {
  it("should return a message", async () => {
    const res = await defaultApi.instance().apiExampleGet();

    const {
      status,
      data: { message },
    } = res;
    expect(status).toBe(200);
    expect(message).toEqual("Hi");
  });
});
