<template>
  <div id="app">
    <nav class="navbar is-dark" role="navigation" aria-label="main navigation">
      <div class="navbar-brand">
        <a class="navbar-item" style="text-align:center;">
          <p style='font-weight: bold;'>EgressosPass</p>
        </a>
        <a
          role="button"
          class="navbar-burger"
          aria-label="menu"
          aria-expanded="false"
        >
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
        </a>
      </div>
    </nav>
    <div class="container is-fluid">
      <Table></Table>
    </div>
    <footer class="footer">
      <div class="content has-text-centered">
        <p>
          <strong>Criado</strong> por
          <a href="https://github.com/ggpereira">Gabriel Gomes Pereira</a>, 
          <a href="https://github.com/odcmello">Otávio da Cruz Mello</a> and 
          <a href="https://github.com/jcamaral">Juliano Colling</a>, 
          Sistemas de Informação, Universidade Federal de Santa Maria, 2019
        </p>
      </div>
    </footer>
  </div>
</template>

<script>
import Table from "./components/Table";

export default {
  name: "app",
  components: {
    Table
  }
};
</script>