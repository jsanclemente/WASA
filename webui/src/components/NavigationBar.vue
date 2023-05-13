


<template>
    <nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
        <div class="position-sticky pt-3 sidebar-sticky">
            <h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
                <span>General</span>
            </h6>
            <ul class="nav flex-column">
                <li class="nav-item">
                    <RouterLink to="/" class="nav-link text-thin">
                        <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#home"/></svg>
                        Home
                    </RouterLink>
                </li>
                <li class="nav-item">
                    <RouterLink to="/link1" class="nav-link text-thin">
                        <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#layout"/></svg>
                        Search
                    </RouterLink>
                </li>
                <li class="nav-item">
                    <RouterLink to="/link2" class="nav-link">
                        <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#key"/></svg>
                        Profile
                    </RouterLink>
                </li>
            </ul>

        </div>
    </nav>
</template>


<style>
</style>