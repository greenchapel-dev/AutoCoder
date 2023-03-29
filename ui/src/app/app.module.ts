import { HttpClientModule } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { IonicStorageModule } from '@ionic/storage-angular';
import { NgxTextDiffModule } from 'ngx-text-diff';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { MaterialModule } from './material/material.module';
import { SettingsComponent } from './settings/settings.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { HomeComponent } from './home/home.component';
import { RoundGaugeComponent } from './components/round-gauge/round-gauge.component';


@NgModule({
  declarations: [
    AppComponent,
    SettingsComponent,
    HomeComponent,
    RoundGaugeComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    MaterialModule,
    HttpClientModule,
    IonicStorageModule.forRoot(),
    NgxTextDiffModule,
    BrowserAnimationsModule,
  ],
  providers: [],
  bootstrap: [AppComponent],
  exports: [MaterialModule]
})
export class AppModule { }
