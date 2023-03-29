import { Component, ElementRef, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { Router } from '@angular/router';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-settings',
  templateUrl: './settings.component.html',
  styleUrls: ['./settings.component.scss']
})
export class SettingsComponent implements OnInit, OnDestroy {

 

  constructor(private router: Router) {

  }

  async ngOnInit(): Promise<void> {
  }

  ngOnDestroy(): void {
  }

  Done() {
    this.router.navigate(['home']);
  }
}
