<template>
  <div>
		<div class="container-fluid">
			<div class="row">
				<div class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
					<div class="main">
						<div class="mb-2">
							<div class="mb-2">
								<h2 class="mt-3 mb-4 text-center">@{{ this.username }}</h2>
								<div class="row text-thin fs-3 text-center">
									<div class="col-4">
										<p> {{ this.nPosts }}</p>
										<p> posts</p>
									</div>
									<div class="col-4">
										<p> {{ this.nFollowers }}</p>
										<p> followers</p>
									</div>
									<div class="col-4">
										<p> {{ this.nFollowing }}</p>
										<p> following</p>
									</div>
								</div>
								<h3 class="text-thin text-center mt-4">Publicaciones</h3>
								<div class="row">

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
			nFollowing: 0
		}
	},

	async created(){
		try {
			console.log(this.id)
			let response = await this.$axios.get("/users/"+ this.id + "/profile")
			this.username = response.data.Username
			this.nPosts = response.data.Nposts
			this.nFollowers = response.data.Nfollowers
			this.nFollowing = response.data.Nfollowing
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