<template>
    <v-data-table
            :headers="headers"
            :items="books"
            :options.sync="options"
            :server-items-length="total"
            :footer-props="{
            showFirstLastPage: true,
            firstIcon: 'mdi-arrow-collapse-left',
            lastIcon: 'mdi-arrow-collapse-right',
            prevIcon: 'mdi-minus',
            nextIcon: 'mdi-plus',
            }"
            hide-action
            item-key="title">
        <template slot="items" slot-scope="props">
            <tr>
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
                loading: true,
                options: {page:1, itemsPerPage:20},
                headers: [{text: 'Title', value: 'title'}, {text: 'Authors', value: 'authors'}],
                books: [],
                total: 500
            }
        },
        async beforeMount() {
            const {page, itemsPerPage} = this.options;
            [this.books,this.total] = await FetchDataFromApi('/books', page, itemsPerPage);
        },
        watch: {
            options: {
                async handler() {
                    const {page, itemsPerPage} = this.options;
                    [this.books,this.total] = await FetchDataFromApi('/books', page, itemsPerPage);
                },
                deep: true
            }
        }
        ,
    }
</script>

<style scoped>

</style>
