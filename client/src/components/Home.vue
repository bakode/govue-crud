<template>
    <div class="contaner">
        <nav class="navbar navbar-light bg-light">
            <a class="navbar-brand" href="#">
                Golang - Vue {CRUD}
            </a>
        </nav>

        <div class="d-flex justify-content-center mt-5 mb-5">
            <div class="card" style="width: 30rem;">
                <div class="card-body">
                    <h5 class="card-title">Add new Example</h5>
                    <form class="p-2">
                        <div class="form-group">
                            <label for="exampleInputTitle">Title</label>
                            <input v-model="title" type="text" class="form-control" aria-describedby="title" placeholder="Enter Title">
                            <small id="title" class="form-text text-muted">Lorem ipsum dolor, sit amet consectetur adipisicing elit.</small>
                        </div>
                        <div class="form-group">
                            <label for="description">Description</label>
                            <textarea v-model="description" class="form-control" id="description" rows="2"></textarea>
                        </div>
                    </form>
                    <button class="btn btn-primary float-right" v-on:click="store()">Save Data</button>
                </div>
            </div>
        </div>

        <h1 class="pl-5">Example list</h1>
        <p class="pl-5">Lorem ipsum dolor sit amet, consectetur adipisicing elit!</p>
        <div class="row p-5">
            <div class="col-12 col-md-4 pb-2" v-for="example in examples" :key="example.Id">
                <div class="card">
                    <div class="card-body">
                        <h5 class="card-title">{{example.Title}}</h5>
                        <p class="card-text">{{example.Description}}</p>
                        <button class="btn btn-danger float-right" v-on:click="destroy(example.Id)">Delete {{example.Id}}</button>
                    </div>
                </div>
            </div>
        </div>

    </div>
</template>

<script>
import axios from 'axios'

export default {
    name: 'Home',
    data () {
        return {
            url: "http://localhost:8080",
            examples: [],
            title: "",
            description: ""
        }
    },
    methods: {
        getAll() {
            axios.get(this.url+'/examples')
                .then(result => result.data)
                .then(items => {
                    this.examples = items.results
                    console.log(this.examples)
                })
        },
        store() {
            if (this.title != "" && this.description != "") {
                var params = new FormData()
                params.append('title', this.title)
                params.append('description', this.description)
                axios.post(this.url+'/examples', params, {
                    headers: {
                        'Content-Type': 'application/x-www-form-urlencoded'
                    }
                })
                .then(function (response) {
                    window.location.reload()
                    console.log(response)
                })
                .catch(function (error) {
                    console.log(error)
                })
            }
        },
        destroy(id) {
            axios({
                method: 'GET',
                url: `${this.url}/examples/${id}/delete`
            })
            window.location.reload()
        }
    },
    mounted() {
        this.getAll()
    },
}
</script>
