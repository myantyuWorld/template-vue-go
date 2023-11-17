import * as validators from "@vuelidate/validators";
import { createI18n } from "vue-i18n";

const messages = {
  ja: {
    validations: {
      required: "{property}は必須です。",
      email: "正しいメールアドレスの形式で入力してください。",
    },
  },
};

const i18n = createI18n({
  legacy: false, // you must set `false`, to use Composition API
  locale: "ja", // set locale
  messages,
});

// // or import { createI18nMessage } from '@vuelidate/validators'
const { createI18nMessage } = validators;

// // Create your i18n message instance. Used for vue-i18n@9
const withI18nMessage = createI18nMessage({ t: i18n.global.t.bind(i18n) });
// for vue-i18n@8
// const withI18nMessage = createI18nMessage({ t: i18n.t.bind(i18n) })

// wrap each validator.
export const required = withI18nMessage(validators.required);
export const email = withI18nMessage(validators.email);
