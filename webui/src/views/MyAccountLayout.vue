<template>
    <div>
        <header class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
			<a class="navbar-brand col-md-3 col-lg-2 me-0 px-3 fs-5 text-thin">WASAPhoto</a>
             <button class="navbar-toggler d-md-none" type="button" data-bs-toggle="collapse" data-bs-target="#sidebarMenu" aria-controls="sidebarMenu" aria-expanded="false" aria-label="Toggle navigation">
                <span class="navbar-toggler-icon"></span>
             </button>
		</header>
        <div class="collapse" :class="sidebarClass" id="sidebarMenu">
            <nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
                <div class="position-sticky pt-3 sidebar-sticky">
                    <h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
                        <span>General</span>
                    </h6>
                    <ul class="nav flex-column">
                        <li class="nav-item text-thin">
                            <RouterLink :to="{ name: 'home'}"  class="nav-link text-thin">
                                <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#home"/></svg>
                                Home
                            </RouterLink>
                        </li>
                        <li class="nav-item">
                            <button type="button" @click="handleOpenModal" class="nav-link text-thin btn border-0" data-bs-toggle="modal" data-bs-target="#modalSearch">
                                <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#search"/></svg>
                                Search
                            </button>
                        </li>
                        <li class="nav-item">
                            <RouterLink :to="{ name: 'profile', params: { id: this.id }, props: true }" class="nav-link text-thin">
                                <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#user"/></svg>
                                Profile
                            </RouterLink>
                        </li>
                        <li class="nav-item">
                            <button type="button" class="nav-link btn text-thin border-0" data-bs-toggle="modal" data-bs-target="#staticBackdrop">
                                <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#upload"/></svg>
                                Upload Photo
                            </button>
                        </li>
                        <li class="nav-item">
                            <RouterLink :to="{ name: 'login'}" class="nav-link text-thin">
                                <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#key"/></svg>
                                Logout
                            </RouterLink>
                        </li>
                    </ul>

                </div>
            </nav>
        </div>
        <router-view :key="$route.fullPath"></router-view>


<!-- Modal for search an user -->
        <div class="modal fade" id="modalSearch" :class="{ 'show': showModalSearch }" data-bs-backdrop="static" data-bs-keyboard="false" tabindex="-1" aria-labelledby="modalSearchLabel" aria-hidden="true">
            <div class="modal-dialog">
                <div class="modal-content">
                    <div class="modal-header">
                        <h1 class="modal-title fs-5 text-thin" id="modalSearchLabel">Search User</h1>
                        <button type="button" @click="handleCloseModal" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                    </div>
                    <div class="modal-body">
                        <ModalSearch @closeModal="handleCloseModal"></ModalSearch>
                    </div>
                </div>
            </div>
        </div>


<!-- Modal for upload photos-->
        <div class="modal fade" id="staticBackdrop" data-bs-backdrop="static" data-bs-keyboard="false" tabindex="-1" aria-labelledby="staticBackdropLabel" aria-hidden="true">
            <div class="modal-dialog modal-dialog-centered">
                <div class="modal-content">
                    <div class="modal-header">
                        <h1 class="modal-title fs-5 text-thin" id="staticBackdropLabel">Upload Photo</h1>
                        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                    </div>
                    <div class="modal-body">
                        <ModalUpload></ModalUpload>
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
            id: parseInt(localStorage.getItem('userId')),
            showModalSearch: false
        };
    },

    methods: {
        handleOpenModal(){
            this.showModalSearch = true
        },

        handleCloseModal(){
            this.showModalSearch = false
        },
    },
    computed: {
        sidebarClass() {
        return {
            'col-md-3': !this.isSmallScreen,
            'col-lg-2': !this.isSmallScreen,
            'd-md-block': !this.isSmallScreen,
            'd-lg-block': !this.isSmallScreen,
            'd-none': this.isSmallScreen,
        };
        },
        isSmallScreen() {
        return window.innerWidth < 992; // Se considera una pantalla pequeña si el ancho de la ventana es menor a 992px (tamaño de pantalla md en Bootstrap)
        },
  }
}

</script>

<style scope>

</style>