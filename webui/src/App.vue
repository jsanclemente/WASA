
<script setup>
import { RouterLink, RouterView } from 'vue-router'
import ErrorMsg from "./components/ErrorMsg.vue"
import router from "@/router"
</script>


<template>
  <div id="app">
    <router-view></router-view> 
  </div>
</template>

<script>
export default {
  name: "App",
  data() {
    return {
		isLoggedIn: false
	};
  },
};
</script>


<style>
  @font-face {
	font-family: 'Roboto-Light';
	src: url('../Roboto/Roboto-Light.ttf') format('truetype');
	font-weight: normal;
	font-style: normal;
	}

	@font-face {
	font-family: 'Roboto-Regular';
	src: url('../Roboto/Roboto-Regular.ttf') format('truetype');
	font-weight: normal;
	font-style: normal;
	}

	@font-face {
	font-family: 'Roboto-Thin';
	src: url('../Roboto/Roboto-Thin.ttf') format('truetype');
	font-weight: normal;
	font-style: normal;
	}

  .font-light {
		font-family: 'Roboto-Light', sans-serif;
	}

	.text-thin {
		font-family: 'Roboto-Thin', sans-serif;
	}

  .text-regular {
    font-family: 'Roboto-Regular', sans-serif;
  }

  
  body {
    font-family: 'Roboto-Light', sans-serif;
    font-size: 16px;
  }
</style>