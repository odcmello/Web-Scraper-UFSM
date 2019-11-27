<template>
  <section>
    <b-table :data="data" :columns="columns"> </b-table>
  </section>
</template>

<script>
    export default {
        name: "Table",
        data() {
            return {
                 data: [
                    { 'id': 1, 'nome': 'Jesse', 'job_title': 'Simmons', 'company': '2016-10-15 13:43:27', 'location': 'Male', 'url':'https://www.google.com', 'curso':'Sistemas de Informação', 'ano_evasao': 2017 },
                    { 'id': 2, 'nome': 'John', 'job_title': 'Simmons', 'company': '2016-10-15 13:43:27', 'location': 'Male', 'url':'https://www.google.com', 'curso':'Sistemas de Informação', 'ano_evasao': 2017 },
                    { 'id': 3, 'nome': 'Tina', 'job_title': 'Simmons', 'company': '2016-10-15 13:43:27', 'location': 'Male', 'url':'https://www.google.com', 'curso':'Sistemas de Informação', 'ano_evasao': 2017 },
                    { 'id': 4, 'nome': 'Clarence','job_title': 'Simmons', 'company': '2016-10-15 13:43:27', 'location': 'Male', 'url':'https://www.google.com', 'curso':'Sistemas de Informação', 'ano_evasao': 2017 },
                    { 'id': 5, 'nome': 'Anne', 'job_title': 'Simmons', 'company': '2016-10-15 13:43:27', 'location': 'Male', 'url':'https://www.google.com', 'curso':'Sistemas de Informação', 'ano_evasao': 2017  }
                ],
                columns: [
                    {
                        field: 'id',
                        label: 'ID',
                        width: '100',
                        numeric: true,
                        searchable: true,
                    },
                    {
                        field: 'nome',
                        label: 'Nome',
                        searchable: true,
                    },
                    {
                        field: 'job_title',
                        label: 'Cargo',
                        searchable: true,
                    },
                    {
                        field: 'company',
                        label: 'empresa',
                        searchable: true,
                    },
                    {
                        field: 'location',
                        label: 'Localização',
                        centered: true,
                        searchable: true,
                    },
                    {
                        field: 'url',
                        label: 'Linkedin',
                    }, 
                    {
                        field: 'curso',
                        label: 'Curso',
                        searchable: true,
                    }, 
                    {
                        field: 'ano_evasao', 
                        label: 'Ano',
                        searchable: true,
                        width: 100
                    }
                ]
            }
        }
    }
</script>

