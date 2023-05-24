<template>
    <div>
        <input
            type="text"
            class="form-control text-thin no-outline"
            placeholder="Type the username of other user"
            v-model="searchQuery"
            @input="searchUsers"
        />
        <ul class="list-group mt-2">
            <button v-for="user in users" :key="user.ID" @click="() => handleClick(user.ID)" data-bs-dismiss="modal" class="list-group-item text-thin bg-hover">
                @{{ user.Username }}
            </button>
        </ul>

        <error-msg v-if="!users" :msg="'No users found'" class="mt-2"></error-msg>
  </div>
</template>

<script>
export default {
    data() {
        return {
            searchQuery: "",
            users: [],
        };
    },

    methods: {
        async searchUsers(){
            try {
                if (this.searchQuery.length === 0) {
                    this.users = []; // 
                    return // No hacer la petición
                }
				const token = localStorage.getItem('token')
                let response = await this.$axios.get("/users", {
                    params: {
                        username: localStorage.getItem("username"),
                        query: this.searchQuery,
                        id: parseInt(localStorage.getItem('userId'))
                    },
                    headers: {
                        Authorization: token
                    }
                });

                this.users = response.data

            }
            catch(error) {
                console.log(error)
            }
        },

        handleClick(userId){
            this.$emit('closeModal')
            this.$router.push(`/myAccount/profile/${userId}`)
            this.searchQuery = ""
            this.users = []
        }
    }


}
</script>

<style scoped>

.no-outline:focus {
  outline: none !important;
  box-shadow: none !important;
  border-color: #000000;
}

.bg-hover:hover {
    background-color: #e0e0e0;
  }

</style>