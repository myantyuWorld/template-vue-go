import { shallowMount } from "@vue/test-utils";
import HelloWorld from "../../src/components/HelloWorld.vue";
import axios from "axios";
import PetRepository from "../../src/apis/petRepository";

// [参考]
// MOCKING AXIOS WITH JEST THROWS ERROR CANNOT READ PROPERTY &#39;INTERCEPTORS OF UNDEFINED
// https://www.appsloveworld.com/jestjs/2/mocking-axios-with-jest-throws-error-cannot-read-property-39interceptors39
jest.mock("axios", () => {
  return {
    create: jest.fn(() => ({
      get: jest.fn(),
      interceptors: {
        request: { use: jest.fn(), eject: jest.fn() },
        response: { use: jest.fn(), eject: jest.fn() },
      },
    })),
  };
});

const BASE_URL = "http://localhost:8080/";
/***
 * コンポーネントの表示に関する単純なテスト
 */
describe("InfoPage.vue", () => {
  // API処理のテスト
  // hogehoge...

  // データをきちんと表示できているテスト
  it("renders props.msg when passed", () => {
    const msg = "new message";
    const wrapper = shallowMount(HelloWorld, {
      props: { msg },
    });
    expect(wrapper.text()).toMatch(msg);
  });
});

/***
 * 任意のメソッドをテストするサンプル
 */
describe("sample test", () => {
  test("two plus two is four", () => {
    expect(1 + 2).toBe(4);
  });
});

/***
 * APIをテストするサンプル
 * [参考] https://vhudyma-blog.eu/3-ways-to-mock-axios-in-jest/
 * |
 * あまり意味ないテスト？（APIのテストで実装されるべきもの？）
 * APIリクエストして、Componentsの状態が期待値となる、ことをテストすることが多い？
 * ｜
 * 例えば、以下の記事のようなケース
 * https://reffect.co.jp/vue/vue-mock-axios/#mock-%E3%81%AE-axios-%E3%81%AE%E3%83%86%E3%82%B9%E3%83%88%E6%96%B9%E6%B3%95
 */
describe("getPetSummary", () => {
  it("API成功テストケース", async () => {
    const petSummary = {
      pet: {
        id: "1",
        name: "なっちゃん",
        age: "",
        sex: "1",
        now_weight: "4000",
        target_weight: "3500",
        birthday: "",
        created_at: "2023-10-30T16:31:33+09:00",
        updated_at: "2023-10-30T16:31:33+09:00",
      },
      dosage_schedules: {
        today: {
          id: "",
          title: "hogehoge",
          date: "2023-11-04T08:56:09.273600841+09:00",
        },
        next: {
          id: "",
          title: "fugafuga",
          date: "2023-11-04T08:56:09.273600882+09:00",
        },
      },
      physical_condition: {
        food: 3,
        sweat: 2,
        toilet: 1,
      },
      memo: {
        id: "",
        title:
          "今月に入って飲む水の量が\n増えた気がする\n次に通院した時に先生に相談する",
        date: "2023-11-04T08:56:09.273600924+09:00",
      },
      schedules: [
        {
          id: 0,
          pet_id: "1",
          title: "トリミング",
          date: "2023-11-04T08:56:09.273601091+09:00",
          location: "ペテモ立川店",
          created_at: "0001-01-01T00:00:00Z",
          updated_at: "0001-01-01T00:00:00Z",
        },
        {
          id: 0,
          pet_id: "1",
          title: "通院",
          date: "2023-11-04T08:56:09.273601132+09:00",
          location: "ホゲホゲ病院",
          created_at: "0001-01-01T00:00:00Z",
          updated_at: "0001-01-01T00:00:00Z",
        },
      ],
    };
    // (axios.get as any).mockResolvedValueOnce(petSummary);
    // when
    const petRepository = new PetRepository();
    const result = petRepository.getPetSummary(1);

    // expect(axios.get).toHaveBeenCalledWith(`${BASE_URL}/pet/1`);
    expect(result).toEqual(petSummary); // resultが`undefined`になるので成功しない
  });
});
