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
    import {FetchDataFromApi} from '@/utils'

    export default {
        name: "books-table",
        data() {
            return {
                headers: [{text: 'Title', value: 'title'}, {text: 'Authors', value: 'authors'}],
                books: []
            }
        },
        async beforeMount() {
            this.books = await FetchDataFromApi('/books')
        }
    }
</script>

<style scoped>

</style>
