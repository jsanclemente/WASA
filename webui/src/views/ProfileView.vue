<template>
	<div>
		  <div class="container-fluid">
			  <div class="row">
				  <div class="col-md-9 ms-sm-auto col-lg-10 px-md-2">
					  <div class="main">
						  <h2 class="mt-3 mb-3 text-center">@{{ this.username }}</h2>
						  <div class="row text-thin fs-4 text-center">
							  <button class="col-md-4 btn border-0 fs-4">
								  <p> {{ this.nPosts }}</p>
								  <p> posts</p>
							  </button>
							  <button class="col-md-4 btn border-0 fs-4"  @click="fetchFollowers" data-bs-toggle="modal" data-bs-target="#followers">
								  <p> {{ this.nFollowers }}</p>
								  <p> followers</p>
							  </button>
							  <div class="col-md-4 btn border-0 fs-4" @click="fetchFollowing" data-bs-toggle="modal" data-bs-target="#following">
								  <p> {{ this.nFollowing }}</p>
								  <p> following</p>
							  </div>
						  </div>
  
						  <SuccessMsg :msg="'Username changed correctly'" v-if="this.usernameChanged"></SuccessMsg>
						  <error-msg :msg="this.messageError" v-if="this.error"></error-msg>
  
						  <div class="row justify-content-center mt-4" v-if="sameUser && !usernameChanged && !error">
							  <div class="col-auto">
								  <input type="text" class="form-control" v-model="newUsername" placeholder="Change your username">
							  </div>
							  <div class="col-auto">
								  <button @click="handleChangeUsername" class="btn btn-dark w-100">Confirm</button>
							  </div>
							</div>

							<!-- Modal Followers -->
							<div class="modal fade" id="followers" data-bs-backdrop="static" data-bs-keyboard="false" tabindex="-1" aria-labelledby="followersLabel" aria-hidden="true">
							<div class="modal-dialog modal-dialog-centered modal-dialog-scrollable">
								<div class="modal-content">
								<div class="modal-header">
									<h1 class="modal-title fs-5" id="followersLabel">Followers</h1>
									<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
								</div>
								<div class="modal-body">
									<ul class="list-group mt-2" v-if="followersList">
										<li v-for="user in followersList" :key="user.ID" class="list-group-item text-thin text-center">
											@{{ user.Username }}
										</li>
									</ul>
									<h1 class="text-thin" v-if="!followersList">You don't have any followers</h1>
								</div>
								</div>
							</div>
							</div>
							<!-- Modal -->

							<!-- Modal Following -->
							<div class="modal fade" id="following" data-bs-backdrop="static" data-bs-keyboard="false" tabindex="-1" aria-labelledby="followingLabel" aria-hidden="true">
							<div class="modal-dialog modal-dialog-centered modal-dialog-scrollable">
								<div class="modal-content">
								<div class="modal-header">
									<h1 class="modal-title fs-5" id="followingLabel">Following</h1>
									<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
								</div>
								<div class="modal-body">
									<ul class="list-group mt-2" v-if="followingList">
										<li v-for="user in followingList" :key="user.ID" class="list-group-item text-thin text-center">
											@{{ user.Username }}
										</li>
									</ul>
									<h1 class="text-thin" v-if="followingList === null">You are not following any user</h1>
								</div>
								</div>
							</div>
							</div>
							<!-- Modal Following -->
  
						  <div class="row justify-content-center mt-4" v-if="!sameUser && !error">
							  <div class="col-auto">
								  <button v-if="!following" @click="handleFollow" :class="{ 'disabled': this.isDisabled }" class="btn btn-dark btn-lg w-100">Follow</button>
								  <button v-if="following" @click="handleFollow" class="btn btn-white btn-lg border-1 border-dark w-100">Unfollow</button>
							  </div>
							  <div class="col-auto">
								  <button v-if="!banned" @click="handleBan" class="btn btn-dark btn-lg w-100">Ban</button>
								  <button v-if="banned" @click="handleBan" class="btn btn-white btn-lg border-1 border-dark w-100">Unban</button>
							  </div>
						  </div>
						  <h3 class="text-thin text-center mt-4">Publicaciones</h3>
						  <div class="row text-center w-75">
							<div class="col-4" v-for="(image, id) in Object.entries(this.images).reverse()" :key="image[0]">
								  <div class="card mt-2 mb-2 ml-2 mr-2">
										<div class="card-image-container">
											<img :src="image[1]" class="card-img-top h-100 border rounded">
										</div>

			
										<div class="dropdown position-absolute top-0 end-0" v-if="sameUser">
											<button class="btn btn-white border-0" type="button" data-bs-toggle="dropdown" aria-expanded="false">
												<i class="material-symbols-outlined">more_horiz</i>
											</button>
											<ul class="dropdown-menu">
												<button class="dropdown-item text-center text-thin" @click="deletePhoto(id)">Delete post</button>
											</ul>
										</div>
								  	</div>
							  </div>
						  </div>
						  
					  </div>
				  </div>
			  </div>
		  </div>
	  </div>
  </template>
  
  <script>
  import ErrorMsg from '../components/ErrorMsg.vue';
  export default {
	components: { ErrorMsg },
  
	  props: ['id'],
  
	  data() {
		  return {
			  username: "",
			  nPosts: 0,
			  nFollowers: 0,
			  nFollowing: 0,
			  posts: [],
			  followers: [],
			  followersList: [],
			  followingList: [],
			  banners: [],
			  sameUser: false,
			  banned: false,
			  following: false,
			  isDisabled: false,
			  error: false,
			  messageError: "",
			  usernameChanged: false,
			  newUsername: "",
			  images: {}
		  }
	  },
  
	  methods: {
		  async handleFollow(){
			  if(this.following){
				  this.unfollowUser()
  
			  }
			  else {
				  this.followUser()
			  }
		  },
  
		  async unfollowUser(){
			  try {
					const token = localStorage.getItem('token')
					let response = await this.$axios.delete("users/" + parseInt(localStorage.getItem('userId')) +  "/following/" + this.id, {
						headers: {
							Authorization: token 
						}
					})
					this.following = false
					this.nFollowers = this.nFollowers - 1
			  }
			  catch(error){
				  console.log(error)
			  }
		  },
		  
		  async followUser(){
			  try {
					const token = localStorage.getItem('token')
					let response = await this.$axios.put("users/" + parseInt(localStorage.getItem('userId')) +  "/following/" + this.id,
					{},
					{
						headers: {
							Authorization: token
						}
					}
					);
					this.following = true
					this.nFollowers = this.nFollowers + 1
			  }
			  catch(error){
				  this.messageError = error.response.data
				  this.showErrorMsg()
			  }
		  },

		  async fetchFollowers(){
			try{
				const token = localStorage.getItem('token')
				const id_user = this.sameUser ? parseInt(localStorage.getItem('userId')): this.id
				let response = await this.$axios.get("users/" + id_user + "/followers", { 
					headers: {
						Authorization: token
					}
				})
				this.followersList = response.data
			}
			catch(error){
				console.log(error)
			}
		  },

		  async fetchFollowing(){
				try{
					const token = localStorage.getItem('token')
					const id_user = this.sameUser ? parseInt(localStorage.getItem('userId')): this.id
					let response = await this.$axios.get("users/" + id_user + "/following", {
						headers: {
							Authorization: token
						}
					})
					this.followingList = response.data
				}
				catch(error){
					console.log(error)
				}
		  },

  
		  async handleBan(){
			  this.banned = !this.banned
			  this.isDisabled = this.banned ? true : false
			  if(this.banned && this.following){
				  this.following = false
				  this.nFollowers = this.nFollowers - 1
				  this.banUser()
			  }
  
			  if(this.banned == true && !this.following){
				  this.banUser()
			  }
			  else {
				  this.unbanUser()
			  }
			  
		  },
  
		  async banUser(){
			  try {
				  const token = localStorage.getItem('token')
				  let response = await this.$axios.put("users/" + parseInt(localStorage.getItem('userId')) +  "/banned/" + this.id,
				  	{},
					{
						headers: {
							Authorization: token
						}
					})
				  
			  }
			  catch(error){
				  console.log(error)
			  }
		  },
  
		  async unbanUser(){
			  try {
				  const token = localStorage.getItem('token')
				  let response = await this.$axios.delete("users/" + parseInt(localStorage.getItem('userId')) +  "/banned/" + this.id, {
					headers: {
						Authorization: token
					}
				  })
			  }
			  catch(error){
				  console.log(error)
			  }
		  },
  
		  showSuccessMsg(){
			  this.usernameChanged = true
			  setTimeout(() => {
				  this.usernameChanged = false;
			  }, 2000); 
		  },
  
		  showErrorMsg(){
			  this.error = true
			  setTimeout(() => {
				  this.error = false;
			  }, 2000); 
		  },
  
		  async handleChangeUsername(){
  
			  const isValidUsername = /^[a-zA-Z0-9_]{4,12}$/.test(this.newUsername)
			  if (!isValidUsername) {
				  this.messageError = "You have to type a valid username (numbers, letters and non-special characters and 4 to 12 characteres length)"
				  this.showErrorMsg()	
				  this.newUsername = ''
				  return
			  }
  
			  try {
				  const token = localStorage.getItem('token')
				  const response = await this.$axios.put(`/users/${this.id}/name`, {
					  username: this.newUsername,
				  }, {
						headers: {
							Authorization: token 
						}
				  });
				  this.username = this.newUsername
				  localStorage.setItem("username",this.username)
				  this.newUsername = ''
  
				  // Show succes msg
				  this.showSuccessMsg()
  
			  }
			  catch(error){
				  this.messageError = error.response.data
				  this.showErrorMsg()
				  this.newUsername = ''
			  }
		  },
  
		  setSameUser(){
			  this.sameUser = this.username === localStorage.getItem("username") ? true : false
		  },
  
		  setFollowing(){
			  const id = parseInt(localStorage.getItem("userId"))
			  if(this.followers === null){
				  return
			  }
			  this.following = this.followers.includes(id) ? true : false
		  },
  
		  setBan(){
			  const id = parseInt(localStorage.getItem("userId"))
			  if(this.banners === null){
				  return
			  }
			  this.banned = this.banners.includes(id) ? true : false
			  if (this.banned){
				  this.isDisabled = true
			  }
		  },
		  
		  async deletePhoto(photoId) {
				if (this.sameUser == true){
					try {
						const token = localStorage.getItem('token')
						console.log(photoId)
						let response = await this.$axios.delete("/photos/" + photoId + "?userId=" + this.id, {
							headers: {
								Authorization: token
							}
						})
						delete this.images[photoId]

						const index = this.posts.indexOf(parseInt(photoId));
						if (index !== -1) {
							this.posts.splice(index, 1);
							this.nPosts-=1
						}

					} 
					catch(error){
						console.log(error)
					}
				}
          }
	  },
  
	  async created(){
		  try {
			  const token = localStorage.getItem('token')
			  let response = await this.$axios.get("/users/"+ this.id + "/profile", {
				headers: {
					Authorization: token
				}
			  })
			  this.username = response.data.Username
			  this.nPosts = response.data.Nposts
			  this.nFollowers = response.data.Nfollowers
			  this.nFollowing = response.data.Nfollowing
			  this.followers = response.data.Followers
			  this.banners = response.data.Banners
			  this.posts = response.data.Posts
			  console.log(this.posts)



			  if (this.posts !== null){
				for (let i = 0; i < this.posts.length; i++) {
				  let image = await this.$axios.get("/photos/"+ this.posts[i] + "/image", {
					headers: {
						Authorization: token
					}
				  })
				  this.images[image.data.ID] = "data:image/jpeg;base64,"+image.data.Image
			  	}
			  }
  
		  
			  // Set variable "sameUser"
			  this.setSameUser()
  
			  if(!this.sameUser){
				  // Set variable "following"
				  this.setFollowing()
				  // Set variable "banned"
				  this.setBan()
			  }
		  }
		  catch(error){
			  console.log(error)
		  }
	  }
  };
  
  
  </script>
  
  <style scoped>
  
  .material-symbols-outlined {
	position: relative;
	padding: 4px;
	border-radius: 50%;
	background-color: #333;
	color: white;
  }

  .card-image-container {
  	height: 250px; /* Ajusta la altura deseada */
	}
  
  </style>