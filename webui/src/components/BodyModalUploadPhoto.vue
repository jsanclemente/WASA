<template>
    <div class="mb-3">
        <div class="d-flex" v-if="!photoUploaded">
            <input type="file" ref="input" @change="previewImage" class="form-control-input">
            <button v-if="image" class="delete-button d-flex justify-content-end" @click="clearImage">
                <i class="material-symbols-outlined">delete</i>
            </button>
        </div>
        <SuccessMsg v-if="photoUploaded" class="mt-2" :msg="'The photo was uploaded correctly'" />
        <ErrorMsg v-if="errorUploading" class="mt-2" msg="Error uploading the photo, please try again" />
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
            image: null,
            file: null,
            photoUploaded: false,
            errorUploading: false,
        }
    },
    methods: {
        previewImage(event) {
            this.file = event.target.files[0];
            if (!this.file) {
                return
            }
            const reader = new FileReader();
            reader.readAsDataURL(this.file);
            reader.onload = (event) => {
                this.image = event.target.result;
            };
        },

        clearImage(){
            this.image = null
            this.$refs.input.value = ''
        },

        showSuccessMsg(){
            this.photoUploaded = true
            setTimeout(() => {
                this.photoUploaded = false;
            }, 2000); 
        },

        showErrorMsg(){
            this.errorUploading = true
            setTimeout(() => {
                this.errorUploading = false;
            }, 2000); 
        },
        

        async uploadPhoto(){
            try {
                const formData = new FormData()
                formData.append('userId', parseInt(localStorage.getItem('userId')))
                formData.append('image', this.file)

                const token = localStorage.getItem('token')
            
                const config = {
                    headers: {
                        'Content-type': 'multipart/form-data',
                        'Authorization': token 
                    }
                };

                let response = await this.$axios.post("/photos", formData,config)
                console.log(response.data)

                //  Clear the input 
                this.clearImage()

                // Show success message
                this.showSuccessMsg()
            }   
            catch(error){
                console.log(error)
                //  Show error message
                this.showErrorMsg()
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