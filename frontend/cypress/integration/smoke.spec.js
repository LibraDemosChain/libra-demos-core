describe('LibraDemosChain Frontend Smoke Test', () => {
  it('Loads the homepage', () => {
    cy.visit('/');
    cy.contains('Altcoins');
  });
});
