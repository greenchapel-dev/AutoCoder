import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RoundGaugeComponent } from './round-gauge.component';

describe('RoundGaugeComponent', () => {
  let component: RoundGaugeComponent;
  let fixture: ComponentFixture<RoundGaugeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ RoundGaugeComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(RoundGaugeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
