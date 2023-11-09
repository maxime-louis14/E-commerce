// AvatarService.js
import { ref } from "vue";

const avatarURL = ref(null);

const setAvatarURL = (url) => {
  avatarURL.value = url;
};

const getAvatarURL = () => {
  return avatarURL.value;
};

export { setAvatarURL, getAvatarURL };
