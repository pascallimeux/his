import axios from 'axios';
var CONFIG = require('../config.json');
var BASE_URL = CONFIG.hisURL;
//const BASE_URL = 'http://192.168.20.83:8000/ocms/v3';

export {getConsentCCVersion};

function getConsentCCVersion() {
  const url = `${BASE_URL}/api/version`;
  console.log('URL: '+url)
  return axios.get(url).then(response => response.data);
}
