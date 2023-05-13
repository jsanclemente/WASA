


<script>

	export default {
		data: function() {
			return {
				posts: [],
				streamEmpty: false
			}
		},

		async created(){
			try {
				const id = parseInt(localStorage.getItem('userId'));
				let response = await this.$axios.get("/feed/"+ id)
				console.log(response.data)
				if(response.data === null){
					this.streamEmpty = true
					return
				}

				this.posts = response.data
				for (let i = 0; i < this.posts.length; i++) {
					const imageData = this.posts[i].Image; // Suponiendo que `post` es un objeto que contiene los datos binarios de la imagen
					const imageUrl = URL.createObjectURL(new Blob([imageData]));
					this.posts[i].Image = imageUrl
					// console.log(imageUrl)					
				}
			}
			catch (error){
				console.log(error)
			}
		},
		methods: {
			async handleLike(postId){
				try {
					const id = parseInt(localStorage.getItem('userId'));
					let response = await this.$axios.put("photos/" + postId  + "/likes/" + id)
					console.log(response.data)
				} catch(error) {
					console.log(error)
				}
			},

			async handleUnlike(postId) {
				try {
					const id = parseInt(localStorage.getItem('userId'));
					let response = await this.$axios.delete("photos/" + postId  + "/likes/" + id)
					console.log(response.data)
				} catch(error) {
					console.log(error)
				}
			},

		}
	}
</script>

<template>
	<div>
		<div class="container-fluid">
			<div class="row">
				<div class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
					<div v-if="!streamEmpty" class="main">
						<Post v-for="item in posts" @like="handleLike" @unlike="handleUnlike" :key="item.ID" :idPost="item.ID" :username="item.Username" :nlikes="item.Nlikes" :comments="item.Ncomments" :date="item.Date" :time="item.Time" :image="item.Image" :likes="item.Likes">
							<!-- Modal -->
							<div class="modal fade" id="exampleModal" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
								<div class="modal-dialog modal-dialog-centered">
									<div class="modal-content">
										<div class="modal-header">
											<h5 class="modal-title" id="exampleModalLabel">Comments</h5>
											<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
										</div>
										<div class="modal-body">
											<BodyModal :idPhoto="idPost"/>
										</div>
									</div>
								</div>
        					</div>
						</Post>		
					</div>
					<div v-if="streamEmpty" class="main">
						<h1 class="text-white">To see your feed you have to follow users or wait them to post any photo!</h1>
					</div>
				</div>  
			</div>
		</div>
	</div>

</template>

<style scope>
	.font-light {
		font-family: 'Roboto-Light', sans-serif;
	}

	.text-thin {
		font-family: 'Roboto-Thin', sans-serif;
	}

	.main {
		display: flex;
		flex-direction: column;
		align-items: center;
		padding: 2rem;
	}
</style>
