import { Injectable } from '@angular/core';
import { Router } from '@angular/router';

export interface MenuIF {
  path: string;
  name: string;
  tool: string;
  menuIcon: string;
  svgIcon: boolean;
}

@Injectable({
  providedIn: 'root'
})
export class MenuService {
  mainMenuItems: MenuIF[] = [];

  constructor(private router: Router) {
  }

  init() {
    this.mainMenuItems = [];
    console.log('configured routes: ', this.router.config);
    this.router.config.forEach(route => {
      if ( route.data === undefined) {
        return;
      }
      else {
        if( route.data.mainMenu) {
          const newManuItem: MenuIF = {
            name: route.data.menuName,
            menuIcon: route.data.menuIcon,
            svgIcon: route.data.svgIcon === true,
            path: route.path!,
            tool: route.data.tool,
          };
          this.mainMenuItems.push(newManuItem);
        }
      }

    });
    console.log('main menus: ', this.mainMenuItems);
    return this.mainMenuItems;
  }
}
