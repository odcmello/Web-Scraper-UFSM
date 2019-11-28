import { http } from './config'; 

export default { 

    requestEgressos: () => {
        return http.get('egressos');
    }

 }