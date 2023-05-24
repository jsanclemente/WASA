<template>
	<div>
		<Comment v-for="comment in listComments" @delete-comment="deleteComment" :key="comment.IdComment" :username="comment.Username" :comment="comment.Comment" :id="comment.IdComment" :idPhoto="this.idPhoto"></Comment>
		<h2 v-if="listEmpty" class="text-thin ">There are no comments yet in this photo</h2>
		<div class="modal-footer input-group">
			<div class="input-group">
				<input class="form-control" ref="commentInput" placeholder="Write your comment" aria-describedby="button-addon2">
				<button class="btn btn-dark" type="button" id="button-addon2" @click="commentPhoto" >Send</button>
			</div>
        </div>
	</div>

</template>

<script>

export default {
	props: {
		idPhoto: Number
	},

	data(){
		return {
			listEmpty: false,
			listComments: []
		}
	},

	methods: {
		deleteComment(commentId) {
			const index = this.listComments.findIndex(comment => comment.IdComment === commentId);
			if (index !== -1) {
				this.listComments.splice(index, 1);
				if (this.listComments.length === 0){
					this.listEmpty = true
				}
			}
			this.$emit('delete-comment', commentId, this.idPhoto)
		},
		// Add a comment to the post
		async commentPhoto(){
			try {

				const comment = this.$refs.commentInput.value
				const token = localStorage.getItem('token')
				let response = await this.$axios.post("/photos/" + this.idPhoto + "/comments", {
					userId: parseInt(localStorage.getItem('userId')),
					comment: comment
				}, {
					headers: {
						Authorization: token
					}
				})
				
				// agrega el comentario a la lista
				this.listComments.push({
					IdComment: response.data,
					Username: localStorage.getItem("username"),
					Comment: comment
				});
				this.listEmpty = false
				// limpia el input
				this.$refs.commentInput.value = "";

				
			}
			catch(error){
				console.log(error)
			}
		}
	},

	async created(){
		try {
			const id = parseInt(localStorage.getItem('userId'))
			const token = localStorage.getItem('token')
			let response = await this.$axios.get("/photos/" + this.idPhoto + "/comments", {
				headers: {
					Authorization: token
				}
			})
			if (response.data !== null){
				this.listComments = response.data
			}
			if (response.data === null) {
				this.listEmpty = true
			}
		}
		catch (error){
			console.log(error)
		}
	}
}
</script>



<style scoped>
	input:focus {
		outline: none;
		border-color: #000000;
		box-shadow: 0 0 0px #000000;
	}
</style>