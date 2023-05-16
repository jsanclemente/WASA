<template>
    <div>
        <header class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
			<a class="navbar-brand col-md-3 col-lg-2 me-0 px-3 fs-5 text-thin">WASAPhoto</a>
		</header>
		<nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
            <div class="position-sticky pt-3 sidebar-sticky">
                <h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
                    <span>General</span>
                </h6>
                <ul class="nav flex-column">
                    <li class="nav-item text-thin">
                        <div to="home"  class="nav-link text-thin">
                            <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#home"/></svg>
                            Home
                        </div>
                    </li>
                    <li class="nav-item">
                        <button type="button" class="nav-link text-thin btn border-0" data-bs-toggle="modal" data-bs-target="#modalSearch">
                            <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#search"/></svg>
                            Search
                        </button>
                    </li>
                    <li class="nav-item">
                        <RouterLink to="profile" class="nav-link text-thin">
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
                </ul>

            </div>
		</nav>
        <router-view></router-view>


<!-- Modal for search an user -->
        <div class="modal fade" id="modalSearch" data-bs-backdrop="static" data-bs-keyboard="false" tabindex="-1" aria-labelledby="modalSearchLabel" aria-hidden="true">
            <div class="modal-dialog">
                <div class="modal-content">
                    <div class="modal-header">
                        <h1 class="modal-title fs-5 text-thin" id="modalSearchLabel">Search User</h1>
                        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                    </div>
                    <div class="modal-body">
                        <ModalSearch></ModalSearch>
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

</script>

<style scope>

</style>