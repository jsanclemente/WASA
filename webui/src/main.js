import {createApp, reactive} from 'vue'
import App from './App.vue'
import { library } from '@fortawesome/fontawesome-svg-core'
import { fas } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import router from './router'
import axios from './services/axios.js';
import ErrorMsg from './components/ErrorMsg.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'
import Post from './components/Post.vue'
import BodyModal from './components/BodyModal.vue'
import Comment from './components/Comment.vue'
import BodyModalUploadPhoto from './components/BodyModalUploadPhoto.vue'
import SuccessMsg from './components/SuccessMsg.vue'
import ModalSearch from './components/ModalSearch.vue'




import './assets/dashboard.css'
import './assets/main.css'



library.add(fas)

const app = createApp(App)
app.config.globalProperties.$axios = axios;
app.component("ErrorMsg", ErrorMsg);
app.component("LoadingSpinner", LoadingSpinner)
app.component("SuccessMsg",SuccessMsg)
app.component("Post",Post)
app.component("BodyModal",BodyModal)
app.component("Comment",Comment)
app.component("ModalUpload", BodyModalUploadPhoto)
app.component("ModalSearch",ModalSearch)
app.component("fa",FontAwesomeIcon)
app.use(router)
app.mount('#app')
