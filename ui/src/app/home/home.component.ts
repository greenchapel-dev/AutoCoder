import { Component, OnDestroy, OnInit } from '@angular/core';


@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit, OnDestroy {

  constructor() { }


  async ngOnInit(): Promise<void> {

  }

  ngOnDestroy(): void {
  }
}
