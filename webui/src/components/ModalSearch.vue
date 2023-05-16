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
            <li v-for="user in users" :key="user.ID" class="list-group-item text-thin bg-hover">
                @{{ user.Username }}
            </li>
        </ul>

        <error-msg v-if="!users" :msg="'No users found'" class="mt-2"></error-msg>
  </div>
</template>

<script>
export default {
    data() {
        return {
            searchQuery: "",
            users: []
        };
    },

    methods: {
        async searchUsers(){
            try {
                if (this.searchQuery.length === 0) {
                    this.users = []; // 
                    return; // Evitar realizar la petición
                }
                let response = await this.$axios.get("/users",{
                    params: {
                        username: this.searchQuery
                    }
                })
                this.users = response.data
            }
            catch(error) {
                console.log(error)
            }
        },
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