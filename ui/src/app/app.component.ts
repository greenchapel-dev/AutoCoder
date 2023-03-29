import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { StorageService } from 'src/app/services/storage.service';
import {Location} from '@angular/common';
import { MatIconRegistry } from "@angular/material/icon";
import { DomSanitizer } from '@angular/platform-browser';
import { MenuIF, MenuService } from './services/menu.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {
  sidenavOpened = false;
  title = 'AutoCoderApp';


  menus: MenuIF[] = [];

  constructor(private router: Router,
              private storageService: StorageService,
              private location: Location,
              private matIconRegistry: MatIconRegistry,
              private domSanitizer: DomSanitizer,
              private menuService: MenuService,
              ) {
    this.router.navigate(['']);

    this.matIconRegistry.addSvgIcon(
      "git-pull-request",
      this.domSanitizer.bypassSecurityTrustResourceUrl("assets/pr.svg")
    );


  }

async ngOnInit(): Promise<void> {
  console.log('init app')
  // start the storage service first
  await this.storageService.init();
  this.menus = await this.menuService.init();

}

  async ngAfterViewInit() {

  }

  toggleSideNav() {
    this.sidenavOpened = !this.sidenavOpened;
  }

  OnGotoSettings() {
    this.router.navigate(['/settings']);
  }

  OnBack() {
    this.location.back()
  }
}
