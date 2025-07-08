func (s *KeeperTestSuite) TestMint() {
    initialSupply := s.bankKeeper.GetSupply(s.ctx, "uldc")
    s.Require().Equal("1000000000000000", initialSupply.Amount.String())
}