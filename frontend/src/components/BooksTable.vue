<template>
    <v-data-table
            :headers="headers"
            :items="books"
            :server-items-length="books.length"
            :items-per-page="5"
            :footer-props="{
            showFirstLastPage: true,
            firstIcon: 'mdi-arrow-collapse-left',
            lastIcon: 'mdi-arrow-collapse-right',
            prevIcon: 'mdi-minus',
            nextIcon: 'mdi-plus'
            }"
            hide-action
            item-key="title">
        <template slot="items" slot-scope="props">
            <tr @click="props.expanded = !props.expanded">
                <td>{{ props.item.title }}</td>
                <td>{{ props.item.authors }}</td>
            </tr>
        </template>
    </v-data-table>

</template>

<script>
    // import Vue from 'vue';
    // import VueResource from 'vue-resource'
    //
    // Vue.use(VueResource);
    // Vue.http.options.emulateJSON = true;
    // const http = Vue.http;

    export default {
        name: "books-table",
        data() {
            return {
                headers: [{text: 'Title', value: 'title'}, {text: 'Authors', value: 'authors'}],
                books: [
                    {title: 'Misery', authors: 'Stephen King', id: 1},
                    {title: 'Dracula', authors: 'Bram Stoker', id: 2}
                ]
            }
        },
        methods: {
            fetchBooks: function () {
                fetch('http://localhost:1323/books',
                    {
                        mode: 'cors', // no-cors, *cors, same-origin
                        headers: {
                            'Access-Control-Allow-Origin': '*'
                        },
                    }
                )
                    .then(response => {
                            if (!response.ok) {
                                console.log('Network response not is ok.');
                            } else {
                                response.json().then(data => {
                                    this.books = data
                                }).catch(error => {
                                        console.log('json error', error);
                                    }
                                )
                            }
                        }
                    ).catch(error => {
                    console.log('There was a problem with your fetch request: ', error.message);
                })
            }
        },
        created() {
            console.log(this.books);
            this.fetchBooks()
        }
    }
</script>

<style scoped>

</style>
