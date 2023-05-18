<template>
  <div>
		<div class="container-fluid">
			<div class="row">
				<div class="col-md-9 ms-sm-auto col-lg-10 px-md-2">
					<div class="main">
						<h2 class="mt-3 mb-4 text-center">@{{ this.username }}</h2>
						<div class="row text-thin fs-4 text-center">
							<div class="col-md-4">
								<p> {{ this.nPosts }}</p>
								<p> posts</p>
							</div>
							<div class="col-md-4">
								<p> {{ this.nFollowers }}</p>
								<p> followers</p>
							</div>
							<div class="col-md-4">
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
							<div class="col-4" v-for="post in this.posts" :key="post.id">
								<div class="card">
									<img src="/logo.png" class="card-img-top" style="width: 100%;">
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
			banners: [],
			sameUser: false,
			banned: false,
			following: false,
			isDisabled: false,
			error: false,
			messageError: "",
			usernameChanged: false,
			newUsername: ""
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
				let response = await this.$axios.delete("users/" + parseInt(localStorage.getItem('userId')) +  "/following/" + this.id)
				this.following = false
				this.nFollowers = this.nFollowers - 1
			}
			catch(error){
				console.log(error)
			}
		},
		
		async followUser(){
			try {
				let response = await this.$axios.put("users/" + parseInt(localStorage.getItem('userId')) +  "/following/" + this.id)
				this.following = true
				this.nFollowers = this.nFollowers + 1
			}
			catch(error){
				this.messageError = error.response.data
				this.showErrorMsg()
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
				let response = await this.$axios.put("users/" + parseInt(localStorage.getItem('userId')) +  "/banned/" + this.id)
				console.log(response.data)
			}
			catch(error){
				console.log(error)
			}
		},

		async unbanUser(){
			try {
				let response = await this.$axios.delete("users/" + parseInt(localStorage.getItem('userId')) +  "/banned/" + this.id)
				console.log(response.data)
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
				const response = await this.$axios.put(`/users/${this.id}/name`, {
					username: this.newUsername,
				});
				console.log(response.data)
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
		}
	},

	async created(){
		try {
			let response = await this.$axios.get("/users/"+ this.id + "/profile")
			this.username = response.data.Username
			this.nPosts = response.data.Nposts
			this.nFollowers = response.data.Nfollowers
			this.nFollowing = response.data.Nfollowing
			this.followers = response.data.Followers
			this.banners = response.data.Banners
			this.posts = response.data.Posts
			// Set variable "sameUser"
			this.setSameUser()

			if(!this.sameUser){
				// Set variable "following"
				this.setFollowing()
				// Set variable "banned"
				this.setBan()
			}

			console.log(response.data)
		}
		catch(error){
			console.log(error)
		}
	}
};


</script>

<style>

</style>