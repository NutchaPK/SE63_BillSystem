import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';

import SignIn from './components/SignIn'
import CreateBill from './components/Bill'

//createHistorytaking
import  create_Patientrights from './components/create_patientrights';


export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
   
    router.registerRoute('/signin', SignIn);
    router.registerRoute('/createbill', CreateBill);

    router.registerRoute('/create_Patientrights', create_Patientrights);
    
  },
});
