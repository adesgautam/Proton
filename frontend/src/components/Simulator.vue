<template>
  <div class="hello">
    <!-- <h1>{{ msg }}</h1> -->
    
    <div class="column is-two-thirds">
    <section class="section">
      <h1 class="title">Rocket Engine Designer</h1>
      <hr>

      
      <b-alert show> Default values reflect the SSME (Space Shuttle Main Engine)</b-alert>
      
      <div>
      <b-container fluid>
        <b-row class="text-center">

          <b-col cols="4">
            <img alt="Engine" src="../assets/engine.png"  width="490">

            <!-- <template>
			  <v-stage :config="configKonva">
			    <v-layer>
			      <v-ellipse :config="ell1"></v-ellipse>
			      <v-ellipse :config="ell2"></v-ellipse>
			      <v-ellipse :config="ell3"></v-ellipse>
			      <v-ellipse :config="ell4"></v-ellipse>
			      <v-line :config="l1"></v-line>
			      <v-line :config="l2"></v-line>
			      <v-line :config="l3"></v-line>
			    </v-layer>
			  </v-stage>
			</template> -->

          </b-col>  
          
          <b-col cols="4">
            <form>
            <div style="text-align: left;">
              <div class="field">
                <label class="label">𝛄, Specific Heat Ratio (Default - 1.2)</label>
                <b-form-input name="gamma" v-model="gamma" v-validate="'required|digits'" class="input" type="text"></b-form-input>
              </div>
              <div class="field" style="margin-top:13px;">
                <label class="label">T<sub>c</sub>, Combustion chamber temp. (Default - 3500 ℃)</label>
                <b-form-input name="tc" v-model="tc" class="input" v-validate="'required|digits'" type="text"></b-form-input>
              </div>

              <div class="field" style="margin-top:13px;">
                <label class="label">M, Mach number (Default - 12)</label>
                <b-form-input name="m" v-model="m" class="input" v-validate="'required|digits'" type="text"></b-form-input>
              </div>

              <div class="field" style="margin-top:13px;">
                <label class="label">P<sub>c</sub>, Combustion chamber pressure (Default - 20 MPa)</label>
                <b-form-input name="pc" v-model="pc" class="input" v-validate="'required|digits'" type="text"></b-form-input>
              </div>

              <div class="field" style="margin-top:13px;">
                <label class="label">P<sub>e</sub>, Exit Pressure (Default - 101 KPa)</label>
                <b-form-input name="pe" v-model="pe" class="input" v-validate="'required|digits'" type="text"></b-form-input>
              </div>

              <div class="field" style="margin-top:13px;">
                <label class="label">ε, Expansion Ratio (Default - 77.5)</label>
                <b-form-input name="epsilon" v-model="epsilon" class="input" v-validate="'required|digits'" type="text"></b-form-input>
              </div>

              <div class="field" style="margin-top:13px;">
                <label class="label">I<sub>sp</sub>, Specific Impulse (Default - 400 sec)</label>
                <b-form-input name="isp" v-model="isp" class="input" v-validate="'required|digits'" type="text"></b-form-input>
              </div>

              <div class="field" style="margin-top:13px;">
                <label class="label">F<sub>th</sub>, Thrust Force (Default - 1.5*1000000 (1.5 MN))</label>
                <b-form-input name="fth" v-model="fth" class="input" v-validate="'required|digits'" type="text"></b-form-input>
              </div>

              <div class="field" style="margin-top:13px;">
                <label class="label">𝛉 - Theta, Nozzle divergence half-angle (Default - 15 degree)</label>
                <b-form-input name="theta" v-model="theta" class="input" v-validate="'required|digits'" type="text"></b-form-input>
              </div>

              <div class="field" style="margin-top:13px;">
                <label class="label">β - Beta, Nozzle convergence half-angle (Default - 60 degree)</label>
                <b-form-input name="beta" v-model="beta" class="input" v-validate="'required|digits'" type="text"></b-form-input>
              </div>

              <div class="field" style="margin-top:13px;">
                <label class="label">𝐠, Acceleration due to gravity (Default - 9.8 m/s)</label>
                <b-form-input name="g" v-model="g" class="input" v-validate="'required|digits'" type="text"></b-form-input>
              </div>

              <div class="field" style="margin-top:13px;">
                <label class="label">σ<sub>c</sub>, Stress on combustion chamber wall (Default - 55 MPa)</label>
                <b-form-input name="sigmac" v-model="sigmac" class="input" v-validate="'required|digits'" type="text"></b-form-input>
              </div>

              </div>
              </form>
            </b-col>
          


          <b-col cols="4">
            <div style="text-align: left;">
            <div><label class="label">Ve (Exhaust velocity) = {{ ve }}</label></div>
            <div style="margin-top:2px;"><label class="label">M<sub>d</sub> (Mass flow rate) = {{ md }}</label></div>
            <div style="margin-top:2px;"><label class="label">A<sub>t</sub> (Throat area) = {{ at }}</label></div>
            <div style="margin-top:2px;"><label class="label">R<sub>t</sub> (Throat radius) = {{ rt }}</label></div>
            <div style="margin-top:2px;"><label class="label">L<sub>dn</sub> (Length of diverging cone) = {{ ldn }}</label></div>
            <div style="margin-top:2px;"><label class="label">A<sub>c</sub> (Area of the inlet at combustion chamber) = {{ ac }}</label></div>
            <div style="margin-top:2px;"><label class="label">R<sub>c</sub> (Radius of the inlet at combustion chamber) = {{ rc }}</label></div>
            <div style="margin-top:2px;"><label class="label">L<sub>cn</sub> (Length of the converging nozzle)</label> = {{ lcn }}</div>
            <div style="margin-top:2px;"><label class="label">L<sub>c</sub> (Length of combustion chamber) = {{ lc }}</label></div>
            <div style="margin-top:2px;"><label class="label">T<sub>wall</sub> (Min. thickness of combustion chamber wall) = {{ twall }}</label></div>
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
// import VueKonva from 'vue-konva'


/* eslint-disable */
Vue.use(VeeValidate)
// Vue.use(VueKonva)

export default {
  name: 'Simulator',

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


<!-- Config values for Kovasjs -->
<!-- configKonva: { width: 500, height: 500 },
      
      ell1: { x: 50, y: 100, radius :{ x : 15, y : 40 },
        stroke: "black", strokeWidth: 1 },

      ell2: { x: 120, y: 100, radius :{ x : 15, y : 40 },
        fill:"#dbdbdb", stroke: "black", strokeWidth:1 },

      ell3: { x: 166, y: 100, radius :{ x : 5, y : 13 },
        fill:"#dbdbdb", stroke: "black", strokeWidth:1 },

      ell4: { x: 380, y: 100, radius :{ x : 15, y : 80 },
         fill:"#dbdbdb", stroke:"black", strokeWidth:1 },

      l1: { points: [50, 60, 120, 60, 166, 87, 380, 19 ], stroke: 'black', strokeWidth:1},
      l2: { points: [50, 140, 120, 140, 166, 112, 380, 181], stroke: 'black', strokeWidth:1},
      l3: { points: [0, 100, 430, 100], stroke: 'black', strokeWidth:1, dash: [5, 5]} -->