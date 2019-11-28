<template>
  <section>
    <b-table :data="egressos" :columns="columns"></b-table>
  </section>
</template>

<script>
import Egressos from "../services/egressos";

export default {
  name: "Table",
  data() {
    return {
      egressos: [],
      columns: [
        {
          field: "id",
          label: "ID",
          width: "100",
          numeric: true,
          searchable: true
        },
        {
          field: "nome",
          label: "Nome",
          searchable: true
        },
        {
          field: "job_title",
          label: "Cargo"
        },
        {
          field: "company",
          label: "empresa",
          searchable: true
        },
        {
          field: "location",
          label: "Localização",
          centered: true,
          searchable: true
        },
        {
          field: "url",
          label: "Linkedin"
        },
        {
          field: "curso",
          label: "Curso",
          searchable: true
        },
        {
          field: "ano_evasao",
          label: "Ano",
          searchable: true,
          width: 100
        }
      ]
    };
  },
  mounted() {
    Egressos.requestEgressos().then(res => {
      this.egressos = res.data;
      console.log(res.data);
    });
  }
};
</script>

