import { Component, NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { SettingsComponent } from './settings/settings.component';


const routes: Routes = [
  {
    path: '',
    redirectTo: '/home',
    pathMatch: 'full'
  },
  {
    path: 'home',
    component: HomeComponent,
    data: {authenticated: false, menuName: 'Home', menuIcon: 'home', mainMenu: true}
  },


  {
    path: 'settings',
    component: SettingsComponent,
    data: {authenticated: true, menuName: 'Settings', menuIcon: 'account_circle', mainMenu: true}
  },

  {
    path: '**',
    redirectTo: ''
  },

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
