<template>
  <div class="card mb-3">
    <div class="card-body d-flex justify-content-start align-items-center">
      <h5 class="card-title mb-0 font-weight-bold">@{{ username }}:</h5>
      <p class="card-text mb-0 text-thin flex-grow-1 d-flex justify-content-center">{{ comment }}</p>
      <button v-if="equals" @click="handleClick" class="delete-button mb-0 d-flex justify-content-end">
        <i class="material-symbols-outlined">delete</i>
      </button> 
    </div>
  </div>
</template>

<script>
export default {
    props: {
        username: String,
        comment: String,
        id: Number,
        idPhoto: Number
    },
    data(){
        return {
            equals: false
        }
    },

    mounted(){
        this.setEquals()
    },
    methods: {
        setEquals(){
            this.equals = this.username === localStorage.getItem("username") ? true : false
        },

        async handleClick(){
            try {
                const token = localStorage.getItem('token')
                let response = await this.$axios.delete("/photos/" + this.idPhoto + "/comments/" + this.id, {
                    headers: {
                        Authorization: token
                    }
                })
                this.$emit('delete-comment',this.id)
            } 
            catch(error){
                console.log(error)
            }
        },
    }
}
</script>

<style>

    .delete-button {
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

    .material-symbols-outlined {
        font-size: 24px;
        transition: all 0.3s ease-in-out;
    }

    .material-symbols-outlined:hover {
        transform: scale(1.2);
    }

</style>