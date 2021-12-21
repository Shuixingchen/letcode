// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.4.16 <0.9.0;

contract HelloWorld{
    string Myname ="lalal";
    function getName() public view returns(string)
    {
        return Myname;
    }
}