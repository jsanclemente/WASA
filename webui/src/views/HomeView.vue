


<script>
import LoadingSpinner from '../components/LoadingSpinner.vue';
	export default {
  components: { LoadingSpinner },
		data: function() {
			return {
				posts: [],
				streamEmpty: false,
				loading: false
			}
		},

		async mounted(){
			try {
				this.loading = true
				const id = parseInt(localStorage.getItem('userId'));
				let response = await this.$axios.get("/feed/"+ id)
				if(response.data === null){
					this.streamEmpty = true
					return
				}

				const array = response.data
				console.log(array)
				this.posts = array.slice().reverse()
				
				
				for (let i = 0; i < this.posts.length; i++) {
					this.posts[i].Image = 'data:image/jpeg;base64,'+this.posts[i].Image
					// console.log(imageUrl)					
				}
				this.loading = false
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
				} catch(error) {
					console.log(error)
				}
			},

			async handleUnlike(postId) {
				try {
					const id = parseInt(localStorage.getItem('userId'));
					let response = await this.$axios.delete("photos/" + postId  + "/likes/" + id)
				} catch(error) {
					console.log(error)
				}
			},

			deleteComment(commentId,photoId){
			}
			
}}
	


</script>

<template>
	<div>
		<div class="container-fluid">
			<div class="row">
				<div class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
					<div v-if="!streamEmpty" class="main">
						<loading-spinner :loading="this.loading"></loading-spinner>
						<Post v-for="item in posts" @delete-comment="deleteComment" @like="handleLike" @unlike="handleUnlike" :key="item.ID" :idPost="item.ID" :username="item.Username" :nlikes="item.Nlikes" :comments="item.Ncomments" :date="item.Date" :time="item.Time" :image="item.Image" :likes="item.Likes">
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
				</div>  
				<div v-if="streamEmpty">
					<h1 class="text-white">To see your feed you have to follow users or wait them to post any photo!</h1>
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