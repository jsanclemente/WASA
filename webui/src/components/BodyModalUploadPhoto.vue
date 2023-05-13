<template>
    <div class="mb-3">
        <div class="d-flex">
            <input type="file" ref="input" @change="previewImage" title="Choose file" class="form-control-file">
            <button v-if="image" class="delete-button d-flex justify-content-end" @click="clearImage">
                <i class="material-symbols-outlined">delete</i>
            </button>
        </div>
        <img :src="image" v-if="image" class="mt-3 img-thumbnail">
        <div class="text-center">
            <button v-if="image" @click="uploadPhoto" class="btn btn-dark mt-3 mx-auto btn-lg w-100">
                Upload
            </button>
        </div>
        
    </div>
</template>

<script>
export default {
    
    data() {
        return {
            image: null
        }
    },
    methods: {
        previewImage(event) {
            const file = event.target.files[0];
            if (!file) {
                return
            }
            const reader = new FileReader();
            reader.readAsDataURL(file);
            reader.onload = (event) => {
                this.image = event.target.result;
            };
        },

        clearImage(){
            this.image = null
            this.$refs.input.value = ''
        },

        async uploadPhoto(){

            const formData = new FormData()
            formData.append('userId', parseInt(localStorage.getItem('userId')))
            formData.append('image', this.image)
            console.log(formData)

            

            // for (const [key, value] of formData.entries()) {
            //     console.log(key, value)
            // }

            const config = {
                headers: {
                    'content-type': 'multipart/form-data'
                }
            };

        
            try {
                let response = await this.$axios.post("/photos", formData,config)
                console.log(response.data)
            }   
            catch(error){
                console.log(error)
            }

        }
        

    }
};

</script>


<style scoped>
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