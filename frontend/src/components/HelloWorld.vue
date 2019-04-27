<template>
  <div class="hello">
    <!-- <h1>{{ msg }}</h1> -->
    
    <div class="column is-two-thirds">
    <section class="section">
      <h1 class="title">Rocket Engine Simulator</h1>
      <hr>

      
      <b-alert show> Default values reflect the SSME (Space Shuttle Main Engine)</b-alert>
      
      <div>
      <b-container fluid>
        <b-row class="text-center">

          <b-col cols="4">
            <img alt="Engine" src="../assets/engine.png"  width="490">
          </b-col>  
          
          <b-col cols="4">
            <form>
            <div style="text-align: left;">
              <div class="field">
                <label class="label">𝛄, Specific Heat Ratio (Default - 1.2)</label>
                <b-form-input name="gamma" v-model="gamma" v-validate="'required|digits'" class="input" type="text"></b-form-input>
              </div>

              <div class="field">
                <label class="label">Tc, Combustion chamber temp. (Default - 3500 ℃)</label>
                <b-form-input name="tc" v-model="tc" class="input" v-validate="'required|digits'" type="text"></b-form-input>
              </div>

              <div class="field">
                <label class="label">M, Mach number (Default - 12)</label>
                <b-form-input name="m" v-model="m" class="input" v-validate="'required|digits'" type="text"></b-form-input>
              </div>

              <div class="field">
                <label class="label">Pc, Combustion chamber pressure (Default - 20 MPa)</label>
                <b-form-input name="pc" v-model="pc" class="input" v-validate="'required|digits'" type="text"></b-form-input>
              </div>

              <div class="field">
                <label class="label">Pe, Exit Pressure (Default - 101 KPa)</label>
                <b-form-input name="pe" v-model="pe" class="input" v-validate="'required|digits'" type="text"></b-form-input>
              </div>

              <div class="field">
                <label class="label">ε, Expansion Ratio (Default - 77.5)</label>
                <b-form-input name="epsilon" v-model="epsilon" class="input" v-validate="'required|digits'" type="text"></b-form-input>
              </div>

              <div class="field">
                <label class="label">Isp, Specific Impulse (Default - 400 sec)</label>
                <b-form-input name="isp" v-model="isp" class="input" v-validate="'required|digits'" type="text"></b-form-input>
              </div>

              <div class="field">
                <label class="label">Fth, Thrust Force (Default - 1.5*1000000 (1.5 MN))</label>
                <b-form-input name="fth" v-model="fth" class="input" v-validate="'required|digits'" type="text"></b-form-input>
              </div>

              <div class="field">
                <label class="label">𝛉 - Theta, Nozzle divergence half-angle (Default - 15 degree)</label>
                <b-form-input name="theta" v-model="theta" class="input" v-validate="'required|digits'" type="text"></b-form-input>
              </div>

              <div class="field">
                <label class="label">β - Beta, Nozzle convergence half-angle (Default - 60 degree)</label>
                <b-form-input name="beta" v-model="beta" class="input" v-validate="'required|digits'" type="text"></b-form-input>
              </div>

              <div class="field">
                <label class="label">𝐠, Acceleration due to gravity (Default - 9.8 m/s)</label>
                <b-form-input name="g" v-model="g" class="input" v-validate="'required|digits'" type="text"></b-form-input>
              </div>

              <div class="field">
                <label class="label">σ<sub>c</sub>, Stress on combustion chamber wall (Default - 55 MPa)</label>
                <b-form-input name="sigmac" v-model="sigmac" class="input" v-validate="'required|digits'" type="text"></b-form-input>
              </div>

              </div>
              </form>
            </b-col>
          


          <b-col cols="4">
            <div style="text-align: left;">
            <div><label class="label">Ve (Exhaust velocity) = {{ ve }}</label></div>
            <div><label class="label">Md (Mass flow rate) = {{ md }}</label></div>
            <div><label class="label">At (Throat area) = {{ at }}</label></div>
            <div><label class="label">Rt (Throat radius) = {{ rt }}</label></div>
            <div><label class="label">Ldn (Length of diverging cone) = {{ ldn }}</label></div>
            <div><label class="label">Ac (Area of the inlet at combustion chamber) = {{ ac }}</label></div>
            <div><label class="label">Rc (Radius of the inlet at combustion chamber) = {{ rc }}</label></div>
            <div><label class="label">Lcn (Length of the converging nozzle)</label> = {{ lcn }}</div>
            <div><label class="label">Lc (Length of combustion chamber) = {{ lc }}</label></div>
            <div><label class="label">Twall (Min. thickness of combustion chamber wall) = {{ twall }}</label></div>
            </div>
          </b-col>
            
          </b-row>
          &nbsp;
          <div>
            <b-button variant="primary" v-on:click="postreq()">Calculate</b-button>  
          </div>
            
        </b-container>
      </div>      
    </section>    
    </div>

  &nbsp;
  <hr>
    
  </div>
</template>


<script>
import axios from 'axios';
import Vue from 'vue'
import VeeValidate from 'vee-validate'
/* eslint-disable */
Vue.use(VeeValidate)

export default {
  name: 'HelloWorld',

  data: function() {
    return {
      ve: "", md: "", at: "", rt: "", ldn: "",
      ac: "", rc: "", lcn: "", lc: "", twall: "",
      gamma : "", tc : "", m : "", pc : "", pe : "",
      epsilon : "", isp : "", fth : "", theta : "", 
      beta : "", g : "", sigmac : ""
    }
  },

  methods: {
    calculate: function() {
      // this.ve = this.g*this.isp
      // this.md = this.fth/this.ve
      
    },  
    postreq: function() {
      var data = { "gamma":parseFloat(this.gamma), "tc":parseFloat(this.tc), "m":parseFloat(this.m), "pc":parseFloat(this.pc), "pe":parseFloat(this.pe),
                "epsilon":parseFloat(this.epsilon), "isp":parseFloat(this.isp), "fth":parseFloat(this.fth), "theta":parseFloat(this.theta), 
                "beta":parseFloat(this.beta), "g":parseFloat(this.g), "sigmac":parseFloat(this.sigmac) }

      /*eslint-disable*/
      console.log(data) 
      /*eslint-enable*/

      axios({ method: "POST", url: "http://127.0.0.1:8090/calc", data: data, headers: {"content-type": "text/plain" } }).then(result => { 
          // this.response = result.data;
          this.ve = result.data['ve'] + ' m/s'
          this.md = result.data['md'] + ' kg/m²'
          this.at = result.data['at'] + ' m²'
          this.rt = result.data['rt'] + ' m'
          this.ldn = result.data['ldn'] + ' m'
          this.ac = result.data['ac'] + ' m²'
          this.rc = result.data['rc'] + ' m'
          this.lcn = result.data['lcn'] + ' m'
          this.lc = result.data['lc'] + ' m'
          this.twall = result.data['twall'] + ' m'
          /*eslint-disable*/
          console.log(result.data) 
          /*eslint-enable*/

        }).catch( error => {
            /*eslint-disable*/
            console.error(error);
            /*eslint-enable*/
        });
    }
  }
}
</script>


<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
eng-details {
  text-align: left;
}
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
