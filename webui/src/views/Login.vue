

<template>
	<div class="container-fluid">
    <div class="row justify-content-center align-items-center vh-100">
      <div class="col-md-6">
        <div class="card border-0 rounded-5 shadow-lg">
		  <h1 class="text-center font-light mt-5 mb-3">WASAPhoto</h1>
          <div class="card-body mb-4 px-5">
            <form @submit.prevent="login">
			  <error-msg v-if="error" :msg=this.errorMsg></error-msg> 
              <div class="mb-3">
                <label for="email" class="form-label fs-4 font-light">Username</label>
                <input type="text" class="form-control text-thin form-control-lg" id="username" placeholder="Write your username" required v-model="username">
              </div>
              <div class="d-grid gap-2">
                <button class="btn btn-dark btn-lg">Log in</button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
	data() {
		return {
			username: '',
			error: false,
			errorMsg: '',
			id: 0,
		}
	},
	methods: {
		async login() {
			const isValidUsername = /^[a-zA-Z0-9_]{4,12}$/.test(this.username);
			if (!isValidUsername) {
				this.error = true
				this.errorMsg = "You have to type a valid username (numbers, letters and non-special characters and 4 to 12 characteres length)"
				return
			}

			try {
				let response = await this.$axios.post("/users", {
					username: this.username,
				});
				this.id = response.data.userId
				localStorage.setItem('userId', String(this.id))
				localStorage.setItem('token',response.data.token)
				localStorage.setItem('username', this.username)
				router.push('/myAccount/home')
			} catch (error) {
				this.error = true
				this.errorMsg = error.response.data.message;
			}
		}
	}
}
		</script>

<style scope>

	input#username:focus {
		outline: none;
		border-color: #000000;
		box-shadow: 0 0 0px #000000;
	}


	

	.font-light {
		font-family: 'Roboto-Light', sans-serif;
	}

	.text-thin {
		font-family: 'Roboto-Thin', sans-serif;
	}

	h1 {
  		font-family: 'Roboto-Regular', sans-serif;
	}

	h2 {
		font-family: 'Roboto-Light', sans-serif;
	}



</style>
