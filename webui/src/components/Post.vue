<template>
    <div class="post mb-1">

        <div class="info">
            <span class="text-regular mb-2">@{{ username }}</span>
            <span class="comments mb-2">{{ date }} - {{  time  }}</span>
        </div>
    
              <!-- Modal -->
							<div class="modal fade" :id="'modal-'+ idPost" tabindex="-1" aria-labelledby="exampleModalLabel" data-bs-backdrop="static" aria-hidden="true">
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
          <!-- End modal -->

        <img :src="image" alt="" class="img-fluid mb-1 border rounded">
        <div class="info">  
            <button class="heart-button">
              <i class="material-symbols-outlined" :class="{ 'active': liked }" @click="handleClick">favorite</i>
              <span class="px-1 likes">{{  !isLiked && liked ? nlikes + 1 : (isLiked && !liked ? nlikes - 1 : nlikes) }} likes</span>
            </button>  
            <button class="comment-button" data-bs-toggle="modal" :data-bs-target="'#modal-' + idPost">
              <i class="material-symbols-outlined">comment</i>
              <span class="comments px-2">
              {{ comments }}   comments
            </span> 
            </button>
        </div>   
    </div>      
  </template>  

  
  <script>
  import BodyModal from "../components/BodyModal.vue"

  export default {
    
    props: {
      idPost: Number,
      nlikes: Number,
      username: String,
      image: String,
      comments: Number,
      username: String,
      date: String,
      time: String,
      likes: Array,
    },

    computed: {
      isLiked() {
        if (this.likes !== null) {
          return this.likes.includes(parseInt(localStorage.getItem('userId')))
        } else {
          return false
        }
      }
    },
    data() {
      return {
        liked: false,
      }
    },
    mounted() {
      this.liked = this.isLiked
    },
    methods: {
      likePost() {
        this.liked = true
        this.$emit('like',this.idPost)
      },
      unlikePost() {
        this.liked = false
        this.$emit('unlike',this.idPost)
      },
      handleClick() {
        if (this.liked) {
          this.unlikePost()
        }
        else {
          this.likePost()
        }
      },

  },

  }
  </script>




  
  <style scoped>

.heart-button {
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: transparent;
  border: none;
  cursor: pointer;
  transition: all 0.2s ease-in-out;
  font-size: 1rem;
  color: #333;
}

.comment-button {
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: transparent;
  border: none;
  cursor: pointer;
  transition: all 0.2s ease-in-out;
  font-size: 1rem;
  color: #333;
}

.heart-button:hover i {
  transform: scale(1.2);
}


  .material-symbols-outlined {
    font-size: 24px;
    transition: all 0.3s ease-in-out;
  }

  .material-symbols-outlined:hover {
    transform: scale(1.2);
  }

  .material-symbols-outlined.active {
    color: red;
  }

  input:focus {
		outline: none;
		border-color: #000000;
		box-shadow: 0 0 0px #000000;
	}


  .post {
    display: flex;
    flex-direction: column;
    align-items: center;
    margin: 1rem;
    border: 2px solid #ccc;
    background-color: #fff;
    box-shadow: 0px 2px 6px rgba(0, 0, 0, 0.3);
    padding: 1rem;
    width: 100%;
  }

  .post img {
    max-width: 500px;
    max-height: 400px;
    width: 100%;
    height: 100%;
    object-fit: contain;
  }

  
  .info {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    width: 100%;
    margin-top: 0.5rem;
  }
  
  .likes {
    font-weight: bold;
    color: #333;
  }
  
  .comments {
    color: #666;
  }
  
  @media only screen and (min-width: 768px) {
    .post {
      max-width: 600px;
      border-radius: 10px;
    }
  }
  </style>
  