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
						<div class="row justify-content-center mt-4" v-if="!sameUser">
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
export default {

	props: ['id'],

	data() {
		return {
			username: "",
			nPosts: 0,
			nFollowers: 0,
			nFollowing: 0,
			posts: [],
			followers: [],
			sameUser: false,
			banned: false,
			following: false,
			isDisabled: false
		}
	},

	methods: {
		handleFollow(){
			this.following = !this.following
			if(this.following === true){
				this.nFollowers = this.nFollowers + 1
			}
			else {
				this.nFollowers = this.nFollowers - 1
			}
			console.log(this.following)
		},

		handleBan(){
			this.banned = !this.banned
			this.isDisabled = this.banned ? true : false
			if(this.banned == true && this.following == true){
				this.following = false
			}
			console.log(this.banned)
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
			this.posts = response.data.Posts
			// Set variable "sameUser"
			this.setSameUser()

			if(!this.sameUser){
				// Set variable "following"
				this.setFollowing()
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